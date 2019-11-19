package database

import "errors"

var (
	ForumIsExist          = errors.New("Forum was created earlier")
	ForumNotFound         = errors.New("Forum not found")
	ForumOrAuthorNotFound = errors.New("Forum or Author not found")
)

const (
	createForumScript = `		
		INSERT INTO forums (slug, title, "user")
		VALUES ($1, $2, (
			SELECT nickname FROM users WHERE nickname = $3
		)) 
		RETURNING "user"`
	getForumScript = `
		SELECT slug, title, "user", posts, threads
		FROM forums
		WHERE slug = $1
	`
)
