package database

import (
	"errors"
	"forum/src/dicts"
	"forum/src/dicts/models"
	"github.com/jackc/pgx"
	"strconv"
	"strings"
)

type ThreadDataManager interface {
	CreateThreadDB(thread *models.Thread) error
	GetThreadDB(slug string) (*models.Thread, error)
	GetForumThreads(slug string, query dicts.QueryParams) ([]*models.Thread, error)
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
func (s service) GetThreadDB(slugOrID string) (*models.Thread, error) {
	sqlRequest := getThreadBySlugScript
	_, err := strconv.Atoi(slugOrID)
	if err != nil {
		sqlRequest = strings.Replace(sqlRequest, "{columnName}", "slug", -1)
	} else {
		sqlRequest = strings.Replace(sqlRequest, "{columnName}", "id", -1)
	}
	thread := &models.Thread{}
	err = s.conn.QueryRow(sqlRequest, slugOrID).Scan(
		&thread.Author,
		&thread.Created,
		&thread.Message,
		&thread.Title,
		&thread.Slug,
		&thread.Forum,
	)
	if err != nil {
		return nil, ThreadNotFound
	}
	return thread, nil
}

//
func (s service) CreateThreadDB(thread *models.Thread) (err error) {
	if thread.Slug != "" {
		_, err := DataManager.GetThreadDB(thread.Slug)
		if err == nil {
			return ThreadIsExist
		}
	}
	if thread == nil {
		return errors.New("Body is not valid ")
	}
	err = s.conn.QueryRow(createThreadScript,
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
		return nil
	case pgxErrNotNull:
		return ForumOrAuthorNotFound //UserNotFound
	case pgxErrForeignKey:
		return ForumOrAuthorNotFound //ForumIsExist
	default:
		return err
	}
}

func (s service) GetForumThreads(slug string, query dicts.QueryParams) ([]*models.Thread, error) {
	var rows *pgx.Rows
	var err error
	threads := make([]*models.Thread, 0)
	sqlRequest := getForumThreadsSinceScript
	sqlRequest = strings.Replace(sqlRequest, "{limit}", query.Limit, -1)
	if query.Desc == "" {
		sqlRequest = strings.Replace(sqlRequest, "DESC", "", -1)
	}
	if query.Since == "" {
		sqlRequest = strings.Replace(sqlRequest, "{sinceQuery}", "", -1)
		rows, err = s.conn.Query(sqlRequest, &slug)
	} else {
		sqlRequest = strings.Replace(sqlRequest, "{sinceQuery}", sinceQuery, -1)
		rows, err = s.conn.Query(sqlRequest, &slug, &query.Since)
	}
	if rows == nil {
		return nil, errors.New("rows is nil")
	}
	if err != nil {
		return nil, ForumNotFound
	}
	defer rows.Close()
	for rows.Next() {
		thread := &models.Thread{}
		err = rows.Scan(
			&thread.Author,
			&thread.Created,
			&thread.Forum,
			&thread.ID,
			&thread.Message,
			&thread.Slug,
			&thread.Title,
			&thread.Votes,
		)
		threads = append(threads, thread)
	}

	if len(threads) == 0 {
		_, err := DataManager.GetForumDB(slug)
		if err != nil {
			return nil, ForumNotFound
		}
	}

	return threads, nil
}

// /thread/{slug_or_id}/vote
func (s service) MakeThreadVoteDB(vote *models.Vote, param string) (*models.Thread, error) {
	var err error

	tx, txErr := s.conn.Begin()
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
		rows, err = s.conn.Query(query, thread.ID, since, limit)
	} else {
		query := queryPostsNoSience[desc][sort]
		rows, err = s.conn.Query(query, thread.ID, limit)
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

	err = s.conn.QueryRow(updateThreadSQL,
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
