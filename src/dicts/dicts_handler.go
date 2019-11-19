package dicts

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func MakeResponse(w http.ResponseWriter, status int, resp interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func MakeErrorUser(s string) string {
	return fmt.Sprintf(`{"message": "Can't find user by nickname: %s"}`, s)
}

func MakeErrorEmail(s string) string {
	return fmt.Sprintf(`{"message": "This email is already registered by user: %s"}`, s)
}

func MakeErrorForum(s string) string {
	return fmt.Sprintf(`{"message": "Can't find forum with slug: %s"}`, s)
}

func MakeErrorThread(s string) string {
	return fmt.Sprintf(`{"message": "Can't find thread by slug: %s"}`, s)
}

func MakeErrorThreadConflict() string {
	return `{"message": "Parent post was created in another thread"}`
}

func MakeErrorThreadID(s string) string {
	return fmt.Sprintf(`{"message": "Can't find thread by slug: %s"}`, s)
}

func MakeErrorPost(s string) string {
	return fmt.Sprintf(`{"message": "Can't find post with id: %s"}`, s)
}

func MakeErrorPostAuthor(s string) string {
	return fmt.Sprintf(`{"message": "Can't find post author by nickname: %s"}`, s)
}
