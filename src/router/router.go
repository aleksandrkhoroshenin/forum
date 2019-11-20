package router

import (
	"forum/src/database"
	"forum/src/dicts/forum"
	"forum/src/dicts/post"
	"forum/src/dicts/thread"
	"forum/src/dicts/user"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/forum/create", forum.CreateForum).Methods("POST")
	router.HandleFunc("/forum/{slug}/create", forum.CreateForumBranch).Methods("POST")
	router.HandleFunc("/forum/{slug}/details", forum.GetBranchDetails).Methods("GET")
	router.HandleFunc("/forum/{slug}/threads", forum.GetBranchThreads).Methods("GET")
	router.HandleFunc("/forum/{slug}/users", forum.GetBranchUsers).Methods("GET")

	router.HandleFunc("/post/{id}/details", post.ChangePostDetails).Methods("POST")
	router.HandleFunc("/post/{id}/details", post.GetPostDetails).Methods("GET")

	router.HandleFunc("/service/clear", database.ClearDB).Methods("POST")
	router.HandleFunc("/service/status", database.GetInformationDB).Methods("GET")

	router.HandleFunc("/user/{nickname}/create", user.CreateUser).Methods("POST")
	router.HandleFunc("/user/{nickname}/profile", user.GetUserInfo).Methods("GET")
	router.HandleFunc("/user/{nickname}/profile", user.ChangeUserInfo).Methods("POST")

	router.HandleFunc("/thread/{slug_or_id}/create", thread.CreateThreadPost).Methods("POST")
	router.HandleFunc("/thread/{slug_or_id}/details", thread.GetThreadDetails).Methods("GET")
	router.HandleFunc("/thread/{slug_or_id}/details", thread.ChangeThreadDetails).Methods("POST")
	router.HandleFunc("/thread/{slug_or_id}/posts", thread.GetPostsFromBranch).Methods("GET")
	router.HandleFunc("/thread/{slug_or_id}/vote", thread.ChangeVoteForBranch).Methods("POST")
	return router
}
