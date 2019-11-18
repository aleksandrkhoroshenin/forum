package forumDataManager

import (
	"forum/src/dicts/models"
	"github.com/jackc/pgx"
)

type ForumDataManager interface {
	CreateForumDB(forum *models.Forum) error
}

type service struct {
	conn *pgx.ConnPool
}

func CreateInstance(conn *pgx.ConnPool) ForumDataManager {
	return service{
		conn: conn,
	}
}

func (s service) CreateForumDB(forum *models.Forum) error {
	err := s.conn.QueryRow(createForumScript,
		&forum.Slug,
		&forum.Title,
		&forum.User).Scan(&forum)
	if err != nil {
		return err
	}
	return nil
}
