package database

import "errors"

const (
	createThreadScript = `
		INSERT
		INTO users ("nickname", "fullname", "email", "about")
		VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING
	`
	getThreadByNicknameOrEmailScript = `
		SELECT "nickname", "fullname", "email", "about"
		FROM users
		WHERE "nickname" = $1 OR "email" = $2
	`
)

var (
	ThreadIsExist  = errors.New("Thread was created earlier")
	ThreadNotFound = errors.New("Thread not found")
)
