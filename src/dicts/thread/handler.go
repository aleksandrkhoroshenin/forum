package thread

import (
	"errors"
	"forum/src/database"
	"forum/src/dicts"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateThreadPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug_or_id"]
	if slug == "" {
		dicts.JsonResponse(w, 400, errors.New("slug_or_id is empty! "))
		return
	}
}

// /thread/{slug_or_id}/details
func GetThreadDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slugOrID := params["slug_or_id"]
	if slugOrID == "" {
		dicts.JsonResponse(w, 400, errors.New("slug_or_id is empty! "))
		return
	}
	thread, err := database.DataManager.GetThreadDB(slugOrID)
	switch err {
	case nil:
		dicts.JsonResponse(w, 200, thread)
	case database.ThreadNotFound:
		dicts.JsonResponse(w, 404, dicts.ErrorThread(slugOrID))
	default:
		dicts.JsonResponse(w, 500, []byte(err.Error()))
	}
}

func ChangeThreadDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug_or_id"]
	if slug == "" {
		dicts.JsonResponse(w, 400, errors.New("slug_or_id is empty! "))
		return
	}
}

func GetPostsFromBranch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug_or_id"]
	if slug == "" {
		dicts.JsonResponse(w, 400, errors.New("slug_or_id is empty! "))
		return
	}
}

func ChangeVoteForBranch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug_or_id"]
	if slug == "" {
		dicts.JsonResponse(w, 400, errors.New("slug_or_id is empty! "))
		return
	}
}
