package database

import "errors"

const (
	createUserSQL = `
		INSERT
		INTO users ("nickname", "fullname", "email", "about")
		VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING
	`
	getUserByNicknameOrEmailSQL = `
		SELECT "nickname", "fullname", "email", "about"
		FROM users
		WHERE "nickname" = $1 OR "email" = $2
	`
	getUserByNickname = `
		SELECT "nickname", "fullname", "email", "about"
		FROM users
		WHERE "nickname" = $1
	`
	getUserSQL = `
		SELECT "nickname", "fullname", "email", "about"
		FROM users
		WHERE "nickname" = $1
	`
	updateUserSQL = `
		UPDATE users
		SET fullname = coalesce(nullif($2, ''), fullname),
			email = coalesce(nullif($3, ''), email),
			about = coalesce(nullif($4, ''), about)
		WHERE "nickname" = $1
		RETURNING nickname, fullname, email, about
	`
)

var (
	UserNotFound       = errors.New("User not found ")
	UserIsExist        = errors.New("User was created earlier ")
	UserUpdateConflict = errors.New("User not updated ")
)
