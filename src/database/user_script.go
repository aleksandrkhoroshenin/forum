package database

import "errors"

const (
	createUserScript = `
		INSERT
		INTO users ("nickname", "fullname", "email", "about")
		VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING
	`
	getUserByNicknameOrEmailScript = `
		SELECT "nickname", "fullname", "email", "about"
		FROM users
		WHERE "nickname" = $1 OR "email" = $2
	`
)

var (
	UserNotFound       = errors.New("User not found")
	UserIsExist        = errors.New("User was created earlier")
	UserUpdateConflict = errors.New("User not updated")
)
