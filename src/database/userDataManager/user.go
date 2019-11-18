package userDataManager

import (
	"forum/src/dicts/models"
	"github.com/jackc/pgx"
)

type UserDataManager interface {
	CreateUserDB(forum *models.User) error
}

type service struct {
	conn *pgx.ConnPool
}

func CreateInstance(conn *pgx.ConnPool) UserDataManager {
	return service{
		conn: conn,
	}
}

func (s service) CreateUserDB(user *models.User) error {
	err := s.conn.QueryRow(createUserScript,
		&user.Nickname,
		&user.Fullname,
		&user.Email,
		&user.About).Scan(&user)
	if err != nil {
		return err
	}
	return nil
}
