package database

import (
	"forum/src/dicts/models"
	"github.com/jackc/pgx"
)

type ForumDataManager interface {
	CreateForumDB(forum *models.Forum) error
	GetForumDB(slug string) (*models.Forum, error)
}

func CreateForumInstance(conn *pgx.ConnPool) ForumDataManager {
	return service{
		conn: conn,
	}
}

func (s service) CreateForumDB(forum *models.Forum) error {
	err := DB.pool.QueryRow(createForumScript,
		&forum.Slug,
		&forum.Title,
		&forum.User).Scan(&forum.User)
	switch ErrorCode(err) {
	case pgxOK:
		return nil
	case pgxErrUnique:
		forum, _ = s.GetForumDB(forum.Slug)
		return ForumIsExist
	case pgxErrNotNull:
		return UserNotFound
	default:
		return err
	}
}

func (s service) GetForumDB(slug string) (*models.Forum, error) {
	f := models.Forum{}

	err := DB.pool.QueryRow(
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
