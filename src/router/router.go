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
	router.POST("/api/v1/forum/create", forum.CreateForum)
	router.POST("/api/v1/forum/{slug}/create", forum.CreateForumBranch)
	router.GET("/api/v1/forum/{slug}/details", forum.GetBranchDetails)
	router.GET("/api/v1/forum/{slug}/threads", forum.GetBranchThreads)
	router.GET("/api/v1/forum/{slug}/users", forum.GetBranchUsers)

	router.POST("/api/v1/post/{id}/details", post.ChangePostDetails)
	router.GET("/api/v1/post/{id}/details", post.GetPostDetails)

	router.POST("/api/v1/service/clear", database.ClearDB)
	router.GET("/api/v1/service/status", database.GetInformationDB)

	router.POST("/api/v1/user/{nickname}/create", user.CreateUser)
	router.GET("/api/v1/user/{nickname}/profile", user.GetUserInfo)
	router.POST("/api/v1/user/{nickname}/profile", user.ChangeUserInfo)

	router.POST("/api/v1/thread/{slug_or_id}/create", thread.CreateThread)
	router.GET("/api/v1/thread/{slug_or_id}/details", thread.CreateThreadBranch)
	router.POST("/api/v1/thread/{slug_or_id}/details", thread.ChangeBranchDetails)
	router.GET("/api/v1/thread/{slug_or_id}/posts", thread.GetPostFromBranch)
	router.POST("/api/v1/thread/{slug_or_id}/vote", thread.ChangeVoteForBranch)
	return router
}
