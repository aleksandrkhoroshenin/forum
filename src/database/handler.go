package database

import (
	"forum/src/database/forumDataManager"
	"forum/src/database/userDataManager"
	"github.com/jackc/pgx"
	"github.com/valyala/fasthttp"
)

type IDataManager struct {
	forumDataManager.ForumDataManager
	userDataManager.UserDataManager
}

var DataManager IDataManager

func CreateDataManagerInstance(conn *pgx.ConnPool) {
	DataManager.ForumDataManager = forumDataManager.CreateInstance(conn)
	DataManager.UserDataManager = userDataManager.CreateInstance(conn)
}

func ClearDB(ctx *fasthttp.RequestCtx) {

}

func GetInformationDB(ctx *fasthttp.RequestCtx) {

}
