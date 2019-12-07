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

	updateThreadSQL = `
		UPDATE threads
		SET title = coalesce(nullif($2, ''), title),
			message = coalesce(nullif($3, ''), message)
		WHERE slug = $1
		RETURNING id, title, author, forum, message, votes, slug, created
	`

	getForumThreadsSinceScript = `
		SELECT author, created, forum, id, message, slug, title, votes
		FROM threads
		WHERE forum = $1 {sinceQuery}
		ORDER BY created DESC
		LIMIT {limit}
	`
	sinceQuery            = ` AND created <= $2::TEXT::TIMESTAMPTZ `
	getThreadBySlugScript = `
		SELECT author, created, forum, id, message, slug, title, votes
		FROM threads
		WHERE {columnName} = $1
	`
	// getThreadPosts
	getPostsSienceDescLimitTreeSQL = `
		SELECT id, author, parent, message, forum, thread, created
		FROM posts
		WHERE thread = $1 AND (path < (SELECT path FROM posts WHERE id = $2::TEXT::INTEGER))
		ORDER BY path DESC
		LIMIT $3::TEXT::INTEGER
	`

	getPostsSienceDescLimitParentTreeSQL = `
		SELECT id, author, parent, message, forum, thread, created
		FROM posts p
		WHERE p.thread = $1 and p.path[1] IN (
			SELECT p2.path[1]
			FROM posts p2
			WHERE p2.thread = $1 AND p2.parent = 0 and p2.path[1] < (SELECT p3.path[1] from posts p3 where p3.id = $2)
			ORDER BY p2.path DESC
			LIMIT $3
		)
		ORDER BY p.path[1] DESC, p.path[2:]
	`

	getPostsSienceDescLimitFlatSQL = `
		SELECT id, author, parent, message, forum, thread, created
		FROM posts
		WHERE thread = $1 AND id < $2::TEXT::INTEGER
		ORDER BY id DESC
		LIMIT $3::TEXT::INTEGER
	`

	getPostsSienceLimitTreeSQL = `
		SELECT id, author, parent, message, forum, thread, created
		FROM posts
		WHERE thread = $1 AND (path > (SELECT path FROM posts WHERE id = $2::TEXT::INTEGER))
		ORDER BY path
		LIMIT $3::TEXT::INTEGER
	`

	getPostsSienceLimitParentTreeSQL = `
		SELECT id, author, parent, message, forum, thread, created
		FROM posts p
		WHERE p.thread = $1 and p.path[1] IN (
			SELECT p2.path[1]
			FROM posts p2
			WHERE p2.thread = $1 AND p2.parent = 0 and p2.path[1] > (SELECT p3.path[1] from posts p3 where p3.id = $2::TEXT::INTEGER)
			ORDER BY p2.path
			LIMIT $3::TEXT::INTEGER
		)
		ORDER BY p.path
	`
	getPostsSienceLimitFlatSQL = `
		SELECT id, author, parent, message, forum, thread, created
		FROM posts
		WHERE thread = $1 AND id > $2::TEXT::INTEGER
		ORDER BY id
		LIMIT $3::TEXT::INTEGER
	`

	// without sience
	getPostsDescLimitTreeSQL = `
		SELECT id, author, parent, message, forum, thread, created
		FROM posts
		WHERE thread = $1 
		ORDER BY path DESC
		LIMIT $2::TEXT::INTEGER
	`
	getPostsDescLimitParentTreeSQL = `
		SELECT id, author, parent, message, forum, thread, created
		FROM posts
		WHERE thread = $1 AND path[1] IN (
			SELECT path[1]
			FROM posts
			WHERE thread = $1
			GROUP BY path[1]
			ORDER BY path[1] DESC
			LIMIT $2::TEXT::INTEGER
		)
		ORDER BY path[1] DESC, path
	`
	getPostsDescLimitFlatSQL = `
		SELECT id, author, parent, message, forum, thread, created
		FROM posts
		WHERE thread = $1
		ORDER BY id DESC
		LIMIT $2::TEXT::INTEGER
	`
	getPostsLimitTreeSQL = `
		SELECT id, author, parent, message, forum, thread, created
		FROM posts
		WHERE thread = $1 
		ORDER BY path
		LIMIT $2::TEXT::INTEGER
	`
	getPostsLimitParentTreeSQL = ` 
		SELECT id, author, parent, message, forum, thread, created
		FROM posts
		WHERE thread = $1 AND path[1] IN (
			SELECT path[1] 
			FROM posts 
			WHERE thread = $1 
			GROUP BY path[1]
			ORDER BY path[1]
			LIMIT $2::TEXT::INTEGER
		)
		ORDER BY path
	`
	getPostsLimitFlatSQL = `
		SELECT id, author, parent, message, forum, thread, created
		FROM posts
		WHERE thread = $1 
		ORDER BY id
		LIMIT $2::TEXT::INTEGER
	`
)

var (
	ThreadIsExist  = errors.New("Thread was created earlier")
	ThreadNotFound = errors.New("Thread not found")
)
