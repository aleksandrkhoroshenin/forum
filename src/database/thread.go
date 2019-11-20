package database

import (
	"errors"
	"forum/src/dicts"
	"forum/src/dicts/models"
	"github.com/jackc/pgx"
	"strconv"
	"strings"
)

type ThreadDataManager interface {
	CreateThreadDB(thread *models.Thread) error
	GetThreadDB(slug string) (*models.Thread, error)
	GetForumThreads(slug string, query dicts.QueryParams) ([]*models.Thread, error)
}

func CreateThreadInstance(conn *pgx.ConnPool) ThreadDataManager {
	return service{
		conn: conn,
	}
}

// /thread/{slug_or_id}/details
func (s service) GetThreadDB(slugOrID string) (*models.Thread, error) {
	sqlRequest := getThreadBySlugScript
	_, err := strconv.Atoi(slugOrID)
	if err != nil {
		sqlRequest = strings.Replace(sqlRequest, "{columnName}", "slug", -1)
	} else {
		sqlRequest = strings.Replace(sqlRequest, "{columnName}", "id", -1)
	}
	thread := &models.Thread{}
	err = s.conn.QueryRow(sqlRequest, slugOrID).Scan(
		&thread.Author,
		&thread.Created,
		&thread.Message,
		&thread.Title,
		&thread.Slug,
		&thread.Forum,
	)
	if err != nil {
		return nil, ThreadNotFound
	}
	return thread, nil
}

//
func (s service) CreateThreadDB(thread *models.Thread) (err error) {
	if thread.Slug != "" {
		_, err := DataManager.GetThreadDB(thread.Slug)
		if err == nil {
			return ThreadIsExist
		}
	}
	if thread == nil {
		return errors.New("Body is not valid ")
	}
	err = s.conn.QueryRow(createThreadScript,
		&thread.Author,
		&thread.Created,
		&thread.Message,
		&thread.Title,
		&thread.Slug,
		&thread.Forum,
	).Scan(
		&thread.Author,
		&thread.Created,
		&thread.Forum,
		&thread.ID,
		&thread.Message,
		&thread.Title,
	)
	switch ErrorCode(err) {
	case pgxOK:
		return nil
	case pgxErrNotNull:
		return ForumOrAuthorNotFound //UserNotFound
	case pgxErrForeignKey:
		return ForumOrAuthorNotFound //ForumIsExist
	default:
		return err
	}
}

func (s service) GetForumThreads(slug string, query dicts.QueryParams) ([]*models.Thread, error) {
	var rows *pgx.Rows
	var err error
	threads := make([]*models.Thread, 0)
	sqlRequest := getForumThreadsSinceScript
	sqlRequest = strings.Replace(sqlRequest, "{limit}", query.Limit, -1)
	if query.Desc == "" {
		sqlRequest = strings.Replace(sqlRequest, "DESC", "", -1)
	}
	if query.Since == "" {
		sqlRequest = strings.Replace(sqlRequest, "{sinceQuery}", "", -1)
		rows, err = s.conn.Query(sqlRequest, &slug)
	} else {
		sqlRequest = strings.Replace(sqlRequest, "{sinceQuery}", sinceQuery, -1)
		rows, err = s.conn.Query(sqlRequest, &slug, &query.Since)
	}
	if rows == nil {
		return nil, errors.New("rows is nil")
	}
	if err != nil {
		return nil, ForumNotFound
	}
	defer rows.Close()
	for rows.Next() {
		thread := &models.Thread{}
		err = rows.Scan(
			&thread.Author,
			&thread.Created,
			&thread.Forum,
			&thread.ID,
			&thread.Message,
			&thread.Slug,
			&thread.Title,
			&thread.Votes,
		)
		threads = append(threads, thread)
	}

	if len(threads) == 0 {
		_, err := DataManager.GetForumDB(slug)
		if err != nil {
			return nil, ForumNotFound
		}
	}

	return threads, nil
}
