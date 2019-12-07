package database

import (
	"github.com/jackc/pgx"
)

type IDataManager struct {
	ForumDataManager
	UserDataManager
	ThreadDataManager
	PostDataManager
	ServiceDataManager
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
	DataManager.ServiceDataManager = CreateServiceInstance(conn)
}
