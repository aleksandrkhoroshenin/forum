package database

import (
	"errors"
	"github.com/jackc/pgx"
)

// Ошибки БД
const (
	pgxOK            = ""
	pgxErrNotNull    = "23502"
	pgxErrForeignKey = "23503"
	pgxErrUnique     = "23505"
	noRowsInResult   = "no rows in result set"
)

// Ошибки запросов
var (
	PostParentNotFound = errors.New("No parent for thread")
	PostNotFound       = errors.New("Post not found")
)

func ErrorCode(err error) string {
	pgerr, ok := err.(pgx.PgError)
	if !ok {
		return pgxOK
	}
	return pgerr.Code
}
