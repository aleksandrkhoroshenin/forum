package thread

import (
	"errors"
	"forum/src/database"
	"forum/src/dicts"
	"forum/src/dicts/models"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func CreateThreadPost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		dicts.JsonResponse(w, 500, err.Error())
		return
	}
	params := mux.Vars(r)
	slugOrID := params["slug_or_id"]
	if slugOrID == "" {
		dicts.JsonResponse(w, 400, errors.New("slugOrID is empty! "))
		return
	}
	post := &models.Post{}
	err = post.UnmarshalJSON(body)
	if err != nil {
		return
	}
	err = database.DataManager.CreatePostDB(slugOrID, post)
	switch err {
	case nil:
		dicts.JsonResponse(w, 201, post)
	case database.ThreadNotFound:
		dicts.JsonResponse(w, 404, dicts.ErrorThreadID(slugOrID))
	case database.UserNotFound:
		dicts.JsonResponse(w, 404, dicts.ErrorPostAuthor(slugOrID))
	case database.PostParentNotFound:
		dicts.JsonResponse(w, 409, dicts.ErrorThreadConflict())
	default:
		dicts.JsonResponse(w, 500, err.Error())
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
