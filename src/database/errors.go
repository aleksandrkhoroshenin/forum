package database

import (
	"errors"
)

// Ошибки БД
const (
	PgxOK            = ""
	PgxErrNotNull    = "23502"
	PgxErrForeignKey = "23503"
	PgxErrUnique     = "23505"
	NoRowsInResult   = "no rows in result set"
)

// Ошибки запросов
var (
	ForumIsExist          = errors.New("Forum was created earlier")
	ForumNotFound         = errors.New("Forum not found")
	ForumOrAuthorNotFound = errors.New("Forum or Author not found")
	ThreadIsExist         = errors.New("Thread was created earlier")
	ThreadNotFound        = errors.New("Thread not found")
	PostParentNotFound    = errors.New("No parent for thread")
	PostNotFound          = errors.New("Post not found")
)
