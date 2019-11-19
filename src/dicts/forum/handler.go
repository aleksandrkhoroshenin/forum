package forum

import (
	"forum/src/database"
	"forum/src/dicts"
	"forum/src/dicts/models"
	"github.com/valyala/fasthttp"
)

func CreateForum(ctx *fasthttp.RequestCtx) {
	forum := &models.Forum{}
	err := forum.UnmarshalJSON(ctx.Request.Body())

	if err != nil {
		return
	}
	err = database.DataManager.CreateForumDB(forum)
	switch err {
	case nil:
		dicts.MakeResponse(ctx, 201, forum)
	case database.UserNotFound:
		dicts.MakeResponse(ctx, 404, forum)
	case database.ForumIsExist:
		dicts.MakeResponse(ctx, 409, forum)
	default:
		dicts.MakeResponse(ctx, 500, forum)
	}
}

func CreateForumBranch(ctx *fasthttp.RequestCtx) {
	ctx.FormValue("")
}

func GetBranchDetails(ctx *fasthttp.RequestCtx) {
	ctx.FormValue("")
}

func GetBranchThreads(ctx *fasthttp.RequestCtx) {
	ctx.FormValue("")
}

func GetBranchUsers(ctx *fasthttp.RequestCtx) {
	ctx.FormValue("")
}
