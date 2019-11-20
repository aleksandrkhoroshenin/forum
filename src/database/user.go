package database

import (
	"forum/src/dicts"
	"forum/src/dicts/models"
	"github.com/jackc/pgx"
)

type UserDataManager interface {
	CreateUserDB(user *models.User) (err error)
	GetUserDB(nickname string) (user *models.User, err error)
	GetForumUsersDB(slug string, query dicts.QueryParams) ([]*models.User, error)
}

func CreateUserInstance(conn *pgx.ConnPool) UserDataManager {
	return service{
		conn: conn,
	}
}

func (s service) CreateUserDB(user *models.User) (err error) {
	rows, err := s.conn.Exec(
		createUserScript,
		&user.Nickname,
		&user.Fullname,
		&user.Email,
		&user.About,
	)
	if err != nil {
		return err
	}

	if rows.RowsAffected() == 0 { // пользователь уже есть
		user := models.User{}
		err := s.conn.QueryRow(
			getUserByNicknameOrEmailScript, &user.Nickname, &user.Email).Scan(&user)

		if err != nil {
			return err
		}

		return UserIsExist
	}

	return nil
}

func (s service) GetUserDB(nickname string) (*models.User, error) {
	user := &models.User{}
	err := s.conn.QueryRow(
		getUserByNicknameScript, &nickname).Scan(
		&user.Nickname,
		&user.Fullname,
		&user.Email,
		&user.About,
	)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s service) GetForumUsersDB(slug string, query dicts.QueryParams) ([]*models.User, error) {
	users := make([]*models.User, 0)
	_, err := s.conn.Query(getForumUsersSinceScript, slug, query.Desc, query.Limit)
	return users, err
}
