package database

const (
	getPostScript = `
		SELECT id, author, message, forum, thread, created, "isEdited", parent
		FROM posts 
		WHERE id = $1
	`
	updatePostScript = `
		UPDATE posts 
		SET message = COALESCE($2, message), "isEdited" = ($2 IS NOT NULL AND $2 <> message) 
		WHERE id = $1 
		RETURNING author::text, created, forum, "isEdited", thread, message, parent
	`
)
