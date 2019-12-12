package database

import (
	"forum/src/dicts/models"
	"github.com/jackc/pgx"
)

const (
	getStatusSQL = `
		SELECT 
		(SELECT COUNT(*) FROM users) AS users,
		(SELECT COUNT(*) FROM forums) AS forums,
		(SELECT COUNT(*) FROM posts) AS posts,
		(SELECT COALESCE(SUM(threads), 0) FROM forums WHERE threads > 0) AS threads
	`
	clearSQL = `
		TRUNCATE users, forums, threads, posts, votes, forum_users;
	`
)

type ServiceDataManager interface {
	GetStatusDB() (*models.Status, error)
	ClearDB()
}

func CreateServiceInstance(conn *pgx.ConnPool) ServiceDataManager {
	return service{
		conn: conn,
	}
}

// /service/status Получение инфомарции о базе данных
func (s service) GetStatusDB() (*models.Status, error) {
	status := &models.Status{}
	err := DB.pool.QueryRow(
		getStatusSQL,
	).Scan(
		&status.User,
		&status.Forum,
		&status.Post,
		&status.Thread,
	)
	if err != nil {
		return nil, err
	}
	return status, nil
}

func (s service) ClearDB() {
	DB.pool.Exec(clearSQL)
}
