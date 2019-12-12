package database

import (
	"fmt"
	"forum/src/dicts/models"
	"github.com/jackc/pgx"
	"strconv"
	"strings"
	"time"
)

type PostDataManager interface {
	CreatePostDB(posts *models.Posts, param string) (*models.Posts, error)
	GetPostDB(id int) (*models.Post, error)
	UpdatePostDB(postUpdate *models.PostUpdate, id int) (*models.Post, error)
	GetPostFullDB(id int, related []string) (*models.PostFull, error)
}

func CreatePostInstance(conn *pgx.ConnPool) PostDataManager {
	return service{
		conn: conn,
	}
}

func (s service) UpdatePostDB(postUpdate *models.PostUpdate, id int) (*models.Post, error) {
	post, err := s.GetPostDB(id)
	if err != nil {
		return nil, PostNotFound
	}

	if len(postUpdate.Message) == 0 {
		return post, nil
	}

	rows := DB.pool.QueryRow(updatePostSQL, strconv.Itoa(id), &postUpdate.Message)

	err = rows.Scan(
		&post.Author,
		&post.Created,
		&post.Forum,
		&post.IsEdited,
		&post.Thread,
		&post.Message,
		&post.Parent,
	)

	if err == nil {
		return post, nil
	} else if err.Error() == noRowsInResult {
		return nil, PostNotFound
	} else {
		return nil, err
	}
}

// /post/{id}/details Получение информации о ветке обсуждения
func (s service) GetPostDB(id int) (*models.Post, error) {
	post := models.Post{}

	err := DB.pool.QueryRow(
		getPostSQL,
		id,
	).Scan(
		&post.ID,
		&post.Author,
		&post.Message,
		&post.Forum,
		&post.Thread,
		&post.Created,
		&post.IsEdited,
		&post.Parent,
	)

	if err == nil {
		return &post, nil
	} else if err.Error() == noRowsInResult {
		return nil, PostNotFound
	} else {
		return nil, err
	}
}

// /post/{id}/details Получение информации о ветке обсуждения
func (s service) GetPostFullDB(id int, related []string) (*models.PostFull, error) {
	postFull := models.PostFull{}
	var err error
	postFull.Post, err = DataManager.GetPostDB(id)
	if err != nil {
		return nil, err
	}

	for _, model := range related {
		switch model {
		case "thread":
			postFull.Thread, err = DataManager.GetThreadDB(strconv.Itoa(int(postFull.Post.Thread)))
		case "forum":
			postFull.Forum, err = DataManager.GetForumDB(postFull.Post.Forum)
		case "user":
			postFull.Author, err = DataManager.GetUserDB(postFull.Post.Author)
		}

		if err != nil {
			return nil, err
		}
	}

	return &postFull, nil
}

// thread/{slug_or_id}/create Создание новых постов
func (s service) CreatePostDB(posts *models.Posts, param string) (*models.Posts, error) {
	thread, err := s.GetThreadDB(param)
	if err != nil {
		return nil, err
	}

	postsNumber := len(*posts)
	if postsNumber == 0 {
		return posts, nil
	}

	dateTimeTemplate := "2006-01-02 15:04:05"
	created := time.Now().Format(dateTimeTemplate)
	query := strings.Builder{}
	query.WriteString("INSERT INTO posts (author, created, message, thread, parent, forum, path) VALUES ")
	queryBody := "('%s', '%s', '%s', %d, %d, '%s', (SELECT path FROM posts WHERE id = %d) || (SELECT last_value FROM posts_id_seq)),"
	for i, post := range *posts {
		err = s.checkPost(post, thread)
		if err != nil {
			return nil, err
		}

		temp := fmt.Sprintf(queryBody, post.Author, created, post.Message, thread.ID, post.Parent, thread.Forum, post.Parent)
		if i == postsNumber-1 {
			temp = temp[:len(temp)-1]
		}
		query.WriteString(temp)
	}
	query.WriteString("RETURNING author, created, forum, id, message, parent, thread")

	tx, txErr := DB.pool.Begin()
	if txErr != nil {
		return nil, txErr
	}
	defer tx.Rollback()

	rows, err := tx.Query(query.String())
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	insertPosts := models.Posts{}
	for rows.Next() {
		post := models.Post{}
		rows.Scan(
			&post.Author,
			&post.Created,
			&post.Forum,
			&post.ID,
			&post.Message,
			&post.Parent,
			&post.Thread,
		)
		insertPosts = append(insertPosts, &post)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	tx.Exec(`UPDATE forums SET posts = posts + $1 WHERE slug = $2`, len(insertPosts), thread.Forum)
	for _, p := range insertPosts {
		tx.Exec(`INSERT INTO forum_users VALUES ($1, $2) ON CONFLICT DO NOTHING`, p.Author, p.Forum)
	}

	tx.Commit()

	return &insertPosts, nil
}

func (s *service) checkPost(p *models.Post, t *models.Thread) error {
	if s.authorExists(p.Author) {
		return UserNotFound
	}
	if s.parentExitsInOtherThread(p.Parent, t.ID) || s.parentNotExists(p.Parent) {
		return PostParentNotFound
	}
	return nil
}

func (s *service) authorExists(nickname string) bool {
	var user models.User
	err := DB.pool.QueryRow(
		getUserByNickname,
		nickname,
	).Scan(
		&user.Nickname,
		&user.Fullname,
		&user.About,
		&user.Email,
	)

	if err != nil && err.Error() == noRowsInResult {
		return true
	}
	return false
}

const postID = `
	SELECT id
	FROM posts
	WHERE id = $1 AND thread IN (SELECT id FROM threads WHERE thread <> $2)
`

func (s *service) parentExitsInOtherThread(parent int64, threadID int32) bool {
	var t int64
	err := DB.pool.QueryRow(postID, parent, threadID).Scan(&t)

	if err != nil && err.Error() == noRowsInResult {
		return false
	}
	return true
}

func (s *service) parentNotExists(parent int64) bool {
	if parent == 0 {
		return false
	}

	var t int64
	err := DB.pool.QueryRow(`SELECT id FROM posts WHERE id = $1`, parent).Scan(&t)

	if err != nil {
		return true
	}
	return false
}
