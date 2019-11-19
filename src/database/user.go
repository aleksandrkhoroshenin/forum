package database

import (
	"forum/src/dicts/models"
	"github.com/jackc/pgx"
)

type UserDataManager interface {
	CreateUserDB(user *models.User) ([]*models.User, error)
	GetUserDB(user *models.User) error
}

func CreateUserInstance(conn *pgx.ConnPool) UserDataManager {
	return service{
		conn: conn,
	}
}

func (s service) CreateUserDB(user *models.User) (users []*models.User, err error) {
	rows, err := s.conn.Exec(
		createUserScript,
		&user.Nickname,
		&user.Fullname,
		&user.Email,
		&user.About,
	)
	if err != nil {
		return nil, err
	}

	if rows.RowsAffected() == 0 { // пользователь уже есть
		user := models.User{}
		err := s.conn.QueryRow(
			getUserByNicknameOrEmailScript, &user.Nickname, &user.Email).Scan(&user)

		if err != nil {
			return nil, err
		}

		return users, UserIsExist
	}

	return nil, nil
}

func (s service) GetUserDB(user *models.User) error {
	err := s.conn.QueryRow(
		getUserByNicknameOrEmailScript, &user.Nickname, &user.Email).Scan(&user)

	if err != nil {
		return err
	}
	return nil
}
