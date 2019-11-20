package database

import (
	"forum/src/dicts"
	"forum/src/dicts/models"
	"github.com/jackc/pgx"
	"strings"
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
		rows, err := s.conn.Query(
			getUserByNicknameOrEmailScript, &user.Nickname, &user.Email)
		defer rows.Close()
		if err != nil {
			return err
		}
		for rows.Next() {
			user := &models.User{}
			err := rows.Scan(&user.Nickname, &user.Fullname, &user.Email, &user.About)
			if err != nil {
				return err
			}
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
	sqlRequest := getForumUsersSinceScript
	if query.Desc == "" {
		sqlRequest = strings.Replace(sqlRequest, "DESC", "", -1)
	}
	rows, err := s.conn.Query(sqlRequest, slug, query.Limit)
	defer rows.Close()

	if err != nil {
		return nil, ForumNotFound
	}

	users := make([]*models.User, 0)
	for rows.Next() {
		u := models.User{}
		err = rows.Scan(
			&u.Nickname,
			&u.Fullname,
			&u.About,
			&u.Email,
		)
		users = append(users, &u)
	}

	if len(users) == 0 {
		_, err := DataManager.GetForumDB(slug)
		if err != nil {
			return nil, ForumNotFound
		}
	}
	return users, err
}
