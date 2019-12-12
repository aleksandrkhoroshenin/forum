package database

import (
	"forum/src/dicts/models"
	"github.com/jackc/pgx"
	"strconv"
)

type ThreadDataManager interface {
	CreateThreadDB(thread *models.Thread) (*models.Thread, error)
	GetThreadDB(param string) (*models.Thread, error)
	GetForumThreads(slug, limit, since, desc string) (*models.Threads, error)
	GetThreadPostsDB(param, limit, since, sort, desc string) (*models.Posts, error)
	UpdateThreadDB(thread *models.ThreadUpdate, param string) (*models.Thread, error)
	MakeThreadVoteDB(vote *models.Vote, param string) (*models.Thread, error)
}

func CreateThreadInstance(conn *pgx.ConnPool) ThreadDataManager {
	return service{
		conn: conn,
	}
}

// /thread/{slug_or_id}/details
func (s service) GetThreadDB(param string) (*models.Thread, error) {
	var err error
	var thread models.Thread

	if isNumber(param) {
		id, _ := strconv.Atoi(param)
		err = DB.pool.QueryRow(
			getThreadIdSQL,
			id,
		).Scan(
			&thread.ID,
			&thread.Title,
			&thread.Author,
			&thread.Forum,
			&thread.Message,
			&thread.Votes,
			&thread.Slug,
			&thread.Created,
		)
	} else {
		err = DB.pool.QueryRow(
			getThreadSlugSQL,
			param,
		).Scan(
			&thread.ID,
			&thread.Title,
			&thread.Author,
			&thread.Forum,
			&thread.Message,
			&thread.Votes,
			&thread.Slug,
			&thread.Created,
		)
	}

	if err != nil {
		return nil, ThreadNotFound
	}

	return &thread, nil
}

// /forum/{slug}/create Создание ветки
func (s service) CreateThreadDB(thread *models.Thread) (*models.Thread, error) {
	if thread.Slug != "" {
		t, err := DataManager.GetThreadDB(thread.Slug)
		if err == nil {
			return t, ThreadIsExist
		}
	}

	err := DB.pool.QueryRow(createThreadScript,
		&thread.Author,
		&thread.Created,
		&thread.Message,
		&thread.Title,
		&thread.Slug,
		&thread.Forum,
	).Scan(
		&thread.Author,
		&thread.Created,
		&thread.Forum,
		&thread.ID,
		&thread.Message,
		&thread.Title,
	)
	switch ErrorCode(err) {
	case pgxOK:
		return thread, nil
	case pgxErrNotNull:
		return nil, ForumOrAuthorNotFound //UserNotFound
	case pgxErrForeignKey:
		return nil, ForumOrAuthorNotFound //ForumIsExist
	default:
		return nil, err
	}
}

var queryForumWithSience = map[string]string{
	"true":  getForumThreadsDescSinceSQL,
	"false": getForumThreadsSinceSQL,
}

var queryForumNoSience = map[string]string{
	"true":  getForumThreadsDescSQL,
	"false": getForumThreadsSQL,
}

// /forum/{slug}/threads
func (s service) GetForumThreads(slug, limit, since, desc string) (*models.Threads, error) {
	var rows *pgx.Rows
	var err error

	if since != "" {
		query := queryForumWithSience[desc]
		rows, err = DB.pool.Query(query, slug, since, limit)
	} else {
		query := queryForumNoSience[desc]
		rows, err = DB.pool.Query(query, slug, limit)
	}
	defer rows.Close()

	if err != nil {
		return nil, ForumNotFound
	}

	threads := models.Threads{}
	for rows.Next() {
		t := models.Thread{}
		err = rows.Scan(
			&t.Author,
			&t.Created,
			&t.Forum,
			&t.ID,
			&t.Message,
			&t.Slug,
			&t.Title,
			&t.Votes,
		)
		threads = append(threads, &t)
	}

	if len(threads) == 0 {
		_, err := s.GetForumDB(slug)
		if err != nil {
			return nil, ForumNotFound
		}
	}
	return &threads, nil
}

// /thread/{slug_or_id}/vote
func (s service) MakeThreadVoteDB(vote *models.Vote, param string) (*models.Thread, error) {
	var err error

	tx, txErr := DB.pool.Begin()
	if txErr != nil {
		return nil, txErr
	}
	defer tx.Rollback()

	var thread models.Thread
	if isNumber(param) {
		id, _ := strconv.Atoi(param)
		err = tx.QueryRow(`SELECT id, author, created, forum, message, slug, title, votes FROM threads WHERE id = $1`, id).Scan(
			&thread.ID,
			&thread.Author,
			&thread.Created,
			&thread.Forum,
			&thread.Message,
			&thread.Slug,
			&thread.Title,
			&thread.Votes,
		)
	} else {
		err = tx.QueryRow(`SELECT id, author, created, forum, message, slug, title, votes FROM threads WHERE slug = $1`, param).Scan(
			&thread.ID,
			&thread.Author,
			&thread.Created,
			&thread.Forum,
			&thread.Message,
			&thread.Slug,
			&thread.Title,
			&thread.Votes,
		)
	}
	if err != nil {
		return nil, ForumNotFound
	}

	var nick string
	err = tx.QueryRow(`SELECT nickname FROM users WHERE nickname = $1`, vote.Nickname).Scan(&nick)
	if err != nil {
		return nil, UserNotFound
	}

	rows, err := tx.Exec(`UPDATE votes SET voice = $1 WHERE thread = $2 AND nickname = $3;`, vote.Voice, thread.ID, vote.Nickname)
	if rows.RowsAffected() == 0 {
		_, err := tx.Exec(`INSERT INTO votes (nickname, thread, voice) VALUES ($1, $2, $3);`, vote.Nickname, thread.ID, vote.Voice)
		if err != nil {
			return nil, UserNotFound
		}
	}
	err = tx.QueryRow(`SELECT votes FROM threads WHERE id = $1`, thread.ID).Scan(&thread.Votes)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	return &thread, nil
}

func isNumber(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}

var queryPostsWithSience = map[string]map[string]string{
	"true": {
		"tree":        getPostsSienceDescLimitTreeSQL,
		"parent_tree": getPostsSienceDescLimitParentTreeSQL,
		"flat":        getPostsSienceDescLimitFlatSQL,
	},
	"false": {
		"tree":        getPostsSienceLimitTreeSQL,
		"parent_tree": getPostsSienceLimitParentTreeSQL,
		"flat":        getPostsSienceLimitFlatSQL,
	},
}

var queryPostsNoSience = map[string]map[string]string{
	"true": {
		"tree":        getPostsDescLimitTreeSQL,
		"parent_tree": getPostsDescLimitParentTreeSQL,
		"flat":        getPostsDescLimitFlatSQL,
	},
	"false": {
		"tree":        getPostsLimitTreeSQL,
		"parent_tree": getPostsLimitParentTreeSQL,
		"flat":        getPostsLimitFlatSQL,
	},
}

// /thread/{slug_or_id}/posts Сообщения данной ветви обсуждения
func (s service) GetThreadPostsDB(param, limit, since, sort, desc string) (*models.Posts, error) {
	thread, err := s.GetThreadDB(param)
	if err != nil {
		return nil, ForumNotFound
	}

	var rows *pgx.Rows

	if since != "" {
		query := queryPostsWithSience[desc][sort]
		rows, err = DB.pool.Query(query, thread.ID, since, limit)
	} else {
		query := queryPostsNoSience[desc][sort]
		rows, err = DB.pool.Query(query, thread.ID, limit)
	}
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	posts := models.Posts{}
	for rows.Next() {
		post := models.Post{}

		err = rows.Scan(
			&post.ID,
			&post.Author,
			&post.Parent,
			&post.Message,
			&post.Forum,
			&post.Thread,
			&post.Created,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &posts, nil
}

// /thread/{slug_or_id}/details Обновление ветки
func (s service) UpdateThreadDB(thread *models.ThreadUpdate, param string) (*models.Thread, error) {
	threadFound, err := s.GetThreadDB(param)
	if err != nil {
		return nil, PostNotFound
	}

	updatedThread := models.Thread{}

	err = DB.pool.QueryRow(updateThreadSQL,
		&threadFound.Slug,
		&thread.Title,
		&thread.Message,
	).Scan(
		&updatedThread.ID,
		&updatedThread.Title,
		&updatedThread.Author,
		&updatedThread.Forum,
		&updatedThread.Message,
		&updatedThread.Votes,
		&updatedThread.Slug,
		&updatedThread.Created,
	)

	if err != nil {
		return nil, err
	}

	return &updatedThread, nil
}
