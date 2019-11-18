package database

import (
	"forum/src/dicts/models"
	"github.com/jackc/pgx"
	"strconv"
)

type PostDataManager interface {
	GetPostDB(id int) (*models.Post, error)
	UpdatePostDB(postUpdate *models.PostUpdate, id int) (*models.Post, error)
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

	rows := s.conn.QueryRow(updatePostScript, strconv.Itoa(id), &postUpdate.Message)

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

func (s service) GetPostDB(id int) (*models.Post, error) {
	post := models.Post{}

	err := s.conn.QueryRow(
		getPostScript,
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
