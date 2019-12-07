package database

import (
	"forum/src/dicts/models"
	"github.com/jackc/pgx"
	"strconv"
)

type PostDataManager interface {
	CreatePostDB(slugOrID string, post *models.Post) error
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

func (s service) CreatePostDB(slugOrID string, post *models.Post) error {
	//thread, err := s.GetThreadDB(slugOrID)
	//if err != nil {
	//	return err
	//}
	//_, err := s.conn.Exec()
	return nil
}
