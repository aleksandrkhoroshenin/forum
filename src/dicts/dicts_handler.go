package dicts

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

func MakeResponse(ctx *fasthttp.RequestCtx, status int, resp interface{}) {
	ctx.Response.Header.Add("Content-Type", "application/json")
	ctx.SetStatusCode(status)
	if err := json.NewEncoder(ctx).Encode(resp); err != nil {
		println(err)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
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
