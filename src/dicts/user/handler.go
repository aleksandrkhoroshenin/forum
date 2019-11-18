package user

import (
	"encoding/json"
	"forum/src/database"
	"forum/src/dicts/models"
	"github.com/valyala/fasthttp"
	"log"
)

func CreateUser(ctx *fasthttp.RequestCtx) {
	nickname := ctx.FormValue("nickname")

	user := &models.User{}
	err := json.Unmarshal(nickname, user.Nickname)
	if err != nil {
		return
	}
	err = user.UnmarshalJSON(ctx.Request.Body())

	if err != nil {
		return
	}
	if err = database.DataManager.CreateUserDB(user); err != nil {
		log.Println(err)
		return
	}
}

func GetUserInfo(ctx *fasthttp.RequestCtx) {
	ctx.FormValue("")
}

func ChangeUserInfo(ctx *fasthttp.RequestCtx) {
	ctx.FormValue("")
}
