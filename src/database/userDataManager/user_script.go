package userDataManager

const (
	createUserScript = `
		INSERT
		INTO users ("nickname", "fullname", "email", "about")
		VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING
	`
)
