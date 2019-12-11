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
		RETURNING "user"
`
	createForumBranch = `
		INSERT INTO threads (author, created, message, title, slug, forum)
		VALUES ($1, $2, $3, $4, $5, (SELECT slug FROM forums WHERE slug = $6)) 
		RETURNING author, created, forum, id, message, title
	`
	getForumScript = `
		SELECT slug, title, "user", posts, threads
		FROM forums
		WHERE slug = $1
	`
)
const (
	getForumThreadsSinceSQL = `
		SELECT author, created, forum, id, message, slug, title, votes
		FROM threads
		WHERE forum = $1 AND created >= $2::TEXT::TIMESTAMPTZ
		ORDER BY created
		LIMIT $3::TEXT::INTEGER
	`
	getForumThreadsDescSinceSQL = `
		SELECT author, created, forum, id, message, slug, title, votes
		FROM threads
		WHERE forum = $1 AND created <= $2::TEXT::TIMESTAMPTZ
		ORDER BY created DESC
		LIMIT $3::TEXT::INTEGER
	`
	getForumThreadsSQL = `
		SELECT author, created, forum, id, message, slug, title, votes
		FROM threads
		WHERE forum = $1
		ORDER BY created
		LIMIT $2::TEXT::INTEGER
	`
	getForumThreadsDescSQL = `
		SELECT author, created, forum, id, message, slug, title, votes
		FROM threads
		WHERE forum = $1
		ORDER BY created DESC
		LIMIT $2::TEXT::INTEGER
	`
	getForumUsersSienceSQl = `
		SELECT nickname, fullname, about, email
		FROM users
		WHERE nickname IN (
				SELECT forum_user FROM forum_users WHERE forum = $1
			) 
			AND LOWER(nickname) > LOWER($2::TEXT)
		ORDER BY nickname
		LIMIT $3::TEXT::INTEGER
	`
	getForumUsersDescSienceSQl = `
		SELECT nickname, fullname, about, email
		FROM users
		WHERE nickname IN (
				SELECT forum_user FROM forum_users WHERE forum = $1
			) 
			AND LOWER(nickname) < LOWER($2::TEXT)
		ORDER BY nickname DESC
		LIMIT $3::TEXT::INTEGER
	`
	getForumUsersSQl = `
		SELECT nickname, fullname, about, email
		FROM users
		WHERE nickname IN (
				SELECT forum_user FROM forum_users WHERE forum = $1
			)
		ORDER BY nickname
		LIMIT $2::TEXT::INTEGER
	`
	getForumUsersDescSQl = `
		SELECT nickname, fullname, about, email
		FROM users
		WHERE nickname IN (
				SELECT forum_user FROM forum_users WHERE forum = $1
			)
		ORDER BY nickname DESC
		LIMIT $2::TEXT::INTEGER
	`
)
