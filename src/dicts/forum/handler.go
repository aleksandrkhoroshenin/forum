package forum

import (
	"forum/src/database"
	"forum/src/dicts/models"
	"github.com/valyala/fasthttp"
	"log"
)

func CreateForum(ctx *fasthttp.RequestCtx) {
	forum := &models.Forum{}
	err := forum.UnmarshalJSON(ctx.Request.Body())

	if err != nil {
		return
	}
	if err = database.DataManager.CreateForumDB(forum); err != nil {
		log.Println(err)
		return
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
