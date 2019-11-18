package database

import (
	"forum/src/database/forumDataManager"
	"github.com/jackc/pgx"
	"github.com/valyala/fasthttp"
)

type IDataManager struct {
	forumDataManager.ForumDataManager
}

var DataManager IDataManager

func CreateDataManagerInstance(conn *pgx.ConnPool) {
	DataManager.ForumDataManager = forumDataManager.CreateInstance(conn)
}

func ClearDB(ctx *fasthttp.RequestCtx) {

}

func GetInformationDB(ctx *fasthttp.RequestCtx) {

}
