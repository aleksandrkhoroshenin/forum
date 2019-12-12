package database

import (
	"forum/src/dicts/models"
	"github.com/jackc/pgx"
)

type ForumDataManager interface {
	CreateForumDB(forum *models.Forum) (*models.Forum, error)
	GetForumDB(slug string) (*models.Forum, error)
}

func CreateForumInstance(conn *pgx.ConnPool) ForumDataManager {
	return service{
		conn: conn,
	}
}

func (s service) CreateForumDB(f *models.Forum) (*models.Forum, error) {
	err := s.conn.QueryRow(createForumScript,
		&f.Slug,
		&f.Title,
		&f.User).Scan(&f.User)
	switch ErrorCode(err) {
	case pgxOK:
		return f, nil
	case pgxErrUnique:
		forum, _ := s.GetForumDB(f.Slug)
		return forum, ForumIsExist
	case pgxErrNotNull:
		return nil, UserNotFound
	default:
		return nil, err
	}
}

func (s service) GetForumDB(slug string) (*models.Forum, error) {
	f := models.Forum{}

	err := s.conn.QueryRow(
		getForumScript,
		slug,
	).Scan(
		&f.Slug,
		&f.Title,
		&f.User,
		&f.Posts,
		&f.Threads,
	)

	if err != nil {
		return nil, ForumNotFound
	}

	return &f, nil
}
