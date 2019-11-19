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

type ErrorMessage struct {
	Message string `json:"message"`
}

func MakeErrorUser(s string) ErrorMessage {
	return ErrorMessage{
		Message: fmt.Sprintf("Can't find user by nickname: %s", s),
	}
}

func MakeErrorEmail(s string) ErrorMessage {
	return ErrorMessage{
		Message: fmt.Sprintf("This email is already registered by user: %s", s),
	}
}

func MakeErrorForum(s string) ErrorMessage {
	return ErrorMessage{
		Message: fmt.Sprintf("Can't find forum with slug: %s", s),
	}
}

func MakeErrorThread(s string) ErrorMessage {
	return ErrorMessage{
		Message: fmt.Sprintf("Can't find thread by slug: %s", s),
	}
}

func MakeErrorThreadConflict() ErrorMessage {
	return ErrorMessage{
		Message: "Parent post was created in another thread",
	}
}

func MakeErrorThreadID(s string) ErrorMessage {
	return ErrorMessage{
		Message: fmt.Sprintf("Can't find thread by slug: %s", s),
	}
}

func MakeErrorPost(s string) ErrorMessage {
	return ErrorMessage{
		Message: fmt.Sprintf("Can't find post with id: %s", s),
	}
}

func MakeErrorPostAuthor(s string) ErrorMessage {
	return ErrorMessage{
		Message: fmt.Sprintf("Can't find post author by nickname: %s", s),
	}
}
