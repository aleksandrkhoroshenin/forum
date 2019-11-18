package database

import (
	"github.com/jackc/pgx"
	"github.com/valyala/fasthttp"
)

type IDataManager struct {
	ForumDataManager
	UserDataManager
	ThreadDataManager
	PostDataManager
}

var DataManager IDataManager

type service struct {
	conn *pgx.ConnPool
}

func CreateDataManagerInstance(conn *pgx.ConnPool) {
	DataManager.ForumDataManager = CreateForumInstance(conn)
	DataManager.UserDataManager = CreateUserInstance(conn)
	DataManager.ThreadDataManager = CreateThreadInstance(conn)
	DataManager.PostDataManager = CreatePostInstance(conn)
}

func ClearDB(ctx *fasthttp.RequestCtx) {

}

func GetInformationDB(ctx *fasthttp.RequestCtx) {

}
