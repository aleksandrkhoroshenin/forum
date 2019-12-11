package database

import (
	"forum/src/dicts/models"
	"github.com/jackc/pgx"
)

type UserDataManager interface {
	CreateUserDB(u *models.User) (*models.Users, error)
	GetUserDB(nickname string) (*models.User, error)
	UpdateUserDB(user *models.User) error
	GetForumUsersDB(slug, limit, since, desc string) (*models.Users, error)
}

func CreateUserInstance(conn *pgx.ConnPool) UserDataManager {
	return service{
		conn: conn,
	}
}

// /user/{nickname}/create Создание нового пользователя
func (s service) CreateUserDB(u *models.User) (*models.Users, error) {
	rows, err := s.conn.Exec(
		createUserSQL,
		&u.Nickname,
		&u.Fullname,
		&u.Email,
		&u.About,
	)
	if err != nil {
		return nil, err
	}

	if rows.RowsAffected() == 0 { // пользователь уже есть
		users := models.Users{}
		queryRows, err := s.conn.Query(getUserByNicknameOrEmailSQL, &u.Nickname, &u.Email)
		defer queryRows.Close()

		if err != nil {
			return nil, err
		}

		for queryRows.Next() {
			user := models.User{}
			queryRows.Scan(&user.Nickname, &user.Fullname, &user.Email, &user.About)
			users = append(users, &user)
		}
		return &users, UserIsExist
	}

	return nil, nil
}

// /user/{nickname}/profile Получение информации о пользователе
func (s service) GetUserDB(nickname string) (*models.User, error) {
	user := models.User{}

	err := s.conn.QueryRow(getUserSQL, nickname).Scan(
		&user.Nickname,
		&user.Fullname,
		&user.Email,
		&user.About,
	)

	if err != nil {
		return nil, UserNotFound
	}

	return &user, nil
}

// /user/{nickname}/profile Изменение данных о пользователе
func (s service) UpdateUserDB(user *models.User) error {
	err := s.conn.QueryRow(
		updateUserSQL,
		&user.Nickname,
		&user.Fullname,
		&user.Email,
		&user.About,
	).Scan(
		&user.Nickname,
		&user.Fullname,
		&user.Email,
		&user.About,
	)

	if err != nil {
		if ErrorCode(err) != pgxOK {
			return UserUpdateConflict
		}
		return UserNotFound
	}

	return nil
}

var queryForumUserWithSience = map[string]string{
	"true":  getForumUsersDescSienceSQl,
	"false": getForumUsersSienceSQl,
}

var queryForumUserNoSience = map[string]string{
	"true":  getForumUsersDescSQl,
	"false": getForumUsersSQl,
}

func (s service) GetForumUsersDB(slug, limit, since, desc string) (*models.Users, error) {
	var rows *pgx.Rows
	var err error

	if since != "" {
		query := queryForumUserWithSience[desc]
		rows, err = s.conn.Query(query, slug, since, limit)
	} else {
		query := queryForumUserNoSience[desc]
		rows, err = s.conn.Query(query, slug, limit)
	}
	defer rows.Close()

	if err != nil {
		return nil, ForumNotFound
	}

	users := models.Users{}
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
		_, err := s.GetForumDB(slug)
		if err != nil {
			return nil, ForumNotFound
		}
	}
	return &users, nil
}
