package database

import "errors"

const (
	createThreadScript = `
		INSERT INTO threads (author, created, message, title, slug, forum)
		VALUES ($1, $2, $3, $4, $5, (SELECT slug FROM forums WHERE slug = $6)) 
		RETURNING author, created, forum, id, message, title
	`

	getForumThreadsScript = `
		SELECT author, created, forum, id, message, slug, title, votes
		FROM threads
		WHERE forum = $1
		ORDER BY created
		LIMIT $2::TEXT::INTEGER
	`

	getThreadByForumAndCreatedScript = `
		SELECT author, created, forum, id, message, slug, title, votes
		FROM threads
		WHERE forum = $1 AND created >= $2::TEXT::TIMESTAMPTZ
		ORDER BY created
		LIMIT $3::TEXT::INTEGER
	`
)

var (
	ThreadIsExist  = errors.New("Thread was created earlier")
	ThreadNotFound = errors.New("Thread not found")
)
