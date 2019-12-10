package router

import (
	"forum/src/dicts/forum"
	"forum/src/dicts/post"
	"forum/src/dicts/service"
	"forum/src/dicts/thread"
	"forum/src/dicts/user"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/forum/create", forum.CreateForum).Methods("POST")
	router.HandleFunc("/api/forum/{slug}/create", forum.CreateForumBranch).Methods("POST")
	router.HandleFunc("/api/forum/{slug}/details", forum.GetBranchDetails).Methods("GET")
	router.HandleFunc("/api/forum/{slug}/threads", forum.GetBranchThreads).Methods("GET")
	router.HandleFunc("/api/forum/{slug}/users", forum.GetBranchUsers).Methods("GET")

	router.HandleFunc("/api/post/{id}/details", post.ChangePostDetails).Methods("POST")
	router.HandleFunc("/api/post/{id}/details", post.GetPostDetails).Methods("GET")

	router.HandleFunc("/api/service/clear", service.ClearDB).Methods("POST")
	router.HandleFunc("/api/service/status", service.GetInformationDB).Methods("GET")

	router.HandleFunc("/api/user/{nickname}/create", user.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/{nickname}/profile", user.GetUserInfo).Methods("GET")
	router.HandleFunc("/api/user/{nickname}/profile", user.ChangeUserInfo).Methods("POST")

	router.HandleFunc("/api/thread/{slug_or_id}/create", thread.CreateThreadPost).Methods("POST")
	router.HandleFunc("/api/thread/{slug_or_id}/details", thread.GetThreadDetails).Methods("GET")
	router.HandleFunc("/api/thread/{slug_or_id}/details", thread.ChangeThreadDetails).Methods("POST")
	router.HandleFunc("/api/thread/{slug_or_id}/posts", thread.GetPostsFromBranch).Methods("GET")
	router.HandleFunc("/api/thread/{slug_or_id}/vote", thread.ChangeVoteForBranch).Methods("POST")
	return router
}
