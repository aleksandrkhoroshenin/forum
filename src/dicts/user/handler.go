package user

import (
	"forum/src/database"
	"forum/src/dicts/models"
	"github.com/valyala/fasthttp"
	"log"
	"strings"
)

func CreateUser(ctx *fasthttp.RequestCtx) {
	args := strings.Split(string(ctx.Request.RequestURI()), "/")
	if len(args) < 4 {
		// 400
		return
	}

	println(ctx.Request.URI().QueryArgs().Peek("nickname"))
	user := &models.User{}
	err := user.UnmarshalJSON(ctx.Request.Body())
	user.Nickname = args[2]

	if err != nil {
		return
	}
	if users, err := database.DataManager.CreateUserDB(user); err != nil {
		log.Println(users, err)
		return
	}
}

func GetUserInfo(ctx *fasthttp.RequestCtx) {
	args := strings.Split(string(ctx.Request.RequestURI()), "/")
	if len(args) < 4 {
		// 400
		return
	}
	user := &models.User{}
	user.Nickname = args[2]
	if err := database.DataManager.GetUserDB(user); err != nil {
		println(err)
		return
	}
}

func ChangeUserInfo(ctx *fasthttp.RequestCtx) {
	ctx.FormValue("")
}
