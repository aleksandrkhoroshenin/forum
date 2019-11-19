package database

import (
	"forum/src/dicts/models"
	"github.com/jackc/pgx"
)

type ThreadDataManager interface {
	CreateThreadDB(user *models.Thread) ([]*models.Thread, error)
}

func CreateThreadInstance(conn *pgx.ConnPool) ThreadDataManager {
	return service{
		conn: conn,
	}
}

func (s service) CreateThreadDB(thread *models.Thread) (threads []*models.Thread, err error) {

	return nil, nil
}
