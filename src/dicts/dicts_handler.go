package dicts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func JsonResponse(w http.ResponseWriter, status int, resp interface{}) {
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

type QueryParams struct {
	Limit  string
	Offset string
	Desc   string
	Since  string
}

func ParseQueryParams(queryParams url.Values) QueryParams {
	params := QueryParams{
		Limit:  queryParams.Get("limit"),
		Offset: queryParams.Get("offset"),
		Since:  queryParams.Get("since"),
		Desc:   queryParams.Get("desc"),
	}
	if params.Limit == "" {
		params.Limit = "1"
	}
	if params.Desc == "true" {
		params.Desc = "desc"
	} else {
		params.Desc = ""
	}
	return params
}

func ErrorFindUserByNick(s string) ErrorMessage {
	return ErrorMessage{
		Message: fmt.Sprintf("Can't find user by nickname: %s", s),
	}
}

func ErrorEmailIsAlreadyExist(s string) ErrorMessage {
	return ErrorMessage{
		Message: fmt.Sprintf("This email is already registered by user: %s", s),
	}
}

func ErrorFindForumWithSlug(s string) ErrorMessage {
	return ErrorMessage{
		Message: fmt.Sprintf("Can't find forum with slug: %s", s),
	}
}

func ErrorThread(s string) ErrorMessage {
	return ErrorMessage{
		Message: fmt.Sprintf("Can't find thread by slug: %s", s),
	}
}

func ErrorThreadConflict() ErrorMessage {
	return ErrorMessage{
		Message: "Parent post was created in another thread",
	}
}

func ErrorThreadID(s string) ErrorMessage {
	return ErrorMessage{
		Message: fmt.Sprintf("Can't find thread by slug: %s", s),
	}
}

func ErrorPost(s string) ErrorMessage {
	return ErrorMessage{
		Message: fmt.Sprintf("Can't find post with id: %s", s),
	}
}

func ErrorPostAuthor(s string) ErrorMessage {
	return ErrorMessage{
		Message: fmt.Sprintf("Can't find post author by nickname: %s", s),
	}
}
