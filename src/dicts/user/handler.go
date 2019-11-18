package user

import (
	"forum/src/database"
	"forum/src/dicts/models"
	"github.com/valyala/fasthttp"
	"log"
)

func CreateUser(ctx *fasthttp.RequestCtx) {
	//nickname := ctx.FormValue("nickname")
	//err := json.Unmarshal(nickname, user.Nickname)
	//if err != nil {
	//	return
	//}

	user := &models.User{}
	err := user.UnmarshalJSON(ctx.Request.Body())

	if err != nil {
		return
	}
	if users, err := database.DataManager.CreateUserDB(user); err != nil {
		log.Println(users, err)
		return
	}
}

func GetUserInfo(ctx *fasthttp.RequestCtx) {
	ctx.FormValue("")
}

func ChangeUserInfo(ctx *fasthttp.RequestCtx) {
	ctx.FormValue("")
}
