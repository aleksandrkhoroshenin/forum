package router

import (
	"forum/src/database"
	"forum/src/dicts/forum"
	"forum/src/dicts/post"
	"forum/src/dicts/thread"
	"forum/src/dicts/user"
	"github.com/buaazp/fasthttprouter"
)

func CreateRouter() *fasthttprouter.Router {
	router := fasthttprouter.New()
	router.POST("/forum/create", forum.CreateForum)
	router.POST("/forum/{slug}/create", forum.CreateForumBranch)
	router.GET("/forum/{slug}/details", forum.GetBranchDetails)
	router.GET("/forum/{slug}/threads", forum.GetBranchThreads)
	router.GET("/forum/{slug}/users", forum.GetBranchUsers)

	router.POST("/post/{id}/details", post.ChangePostDetails)
	router.GET("/post/{id}/details", post.GetPostDetails)

	router.POST("/service/clear", database.ClearDB)
	router.GET("/service/status", database.GetInformationDB)

	router.POST("/user/{nickname}/create", user.CreateUser)
	router.GET("/user/{nickname}/profile", user.GetUserInfo)
	router.POST("/user/{nickname}/profile", user.ChangeUserInfo)

	router.POST("/thread/{slug_or_id}/create", thread.CreateThread)
	router.GET("/thread/{slug_or_id}/details", thread.CreateThreadBranch)
	router.POST("/thread/{slug_or_id}/details", thread.ChangeBranchDetails)
	router.GET("/thread/{slug_or_id}/posts", thread.GetPostFromBranch)
	router.POST("/thread/{slug_or_id}/vote", thread.ChangeVoteForBranch)
	return router
}
