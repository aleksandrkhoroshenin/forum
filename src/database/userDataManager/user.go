package userDataManager

import (
	"forum/src/dicts/models"
	"github.com/jackc/pgx"
)

type UserDataManager interface {
	CreateUserDB(user *models.User) ([]*models.User, error)
}

type service struct {
	conn *pgx.ConnPool
}

func CreateInstance(conn *pgx.ConnPool) UserDataManager {
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
		queryRows, err := s.conn.Query(getUserByNicknameOrEmailScript, &user.Nickname, &user.Email)
		defer queryRows.Close()

		if err != nil {
			return nil, err
		}

		for queryRows.Next() {
			user := models.User{}
			queryRows.Scan(&user.Nickname, &user.Fullname, &user.Email, &user.About)
			users = append(users, &user)
		}
		return users, UserIsExist
	}

	return nil, nil
}
