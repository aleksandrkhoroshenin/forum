package thread

import (
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

	posts := &models.Posts{}
	err = posts.UnmarshalJSON(body)
	if err != nil {
		return
	}
	res, err := database.DataManager.CreatePostDB(posts, slugOrID)
	switch err {
	case nil:
		dicts.JsonResponse(w, 201, res)
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

// /thread/{slug_or_id}/details
func ChangeThreadDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	param := params["slug_or_id"]

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		dicts.JsonResponse(w, 500, []byte(err.Error()))
		return
	}
	threadUpdate := &models.ThreadUpdate{}
	err = threadUpdate.UnmarshalJSON(body)

	//err = forum.Validate()
	if err != nil {
		dicts.JsonResponse(w, 500, []byte(err.Error()))
		return
	}

	result, err := database.DataManager.UpdateThreadDB(threadUpdate, param)

	switch err {
	case nil:
		resp, _ := result.MarshalJSON()
		dicts.JsonResponse(w, 200, resp)
	case database.PostNotFound:
		dicts.JsonResponse(w, 404, dicts.ErrorThread(param))
	default:
		dicts.JsonResponse(w, 500, err.Error())
	}
}

// /thread/{slug_or_id}/posts
func GetPostsFromBranch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	param := params["slug_or_id"]
	queryParams := r.URL.Query()
	var limit, since, sort, desc string
	if limit = queryParams.Get("limit"); limit == "" {
		limit = "1"
	}
	since = queryParams.Get("since")
	if sort = queryParams.Get("sort"); sort == "" {
		sort = "flat"
	}
	if desc = queryParams.Get("desc"); desc == "" {
		desc = "false"
	}
	result, err := database.DataManager.GetThreadPostsDB(param, limit, since, sort, desc)

	switch err {
	case nil:
		resp, _ := result.MarshalJSON()
		dicts.JsonResponse(w, 200, resp)
	case database.ForumNotFound:
		dicts.JsonResponse(w, 404, dicts.ErrorThread(param))
	default:
		dicts.JsonResponse(w, 500, err.Error())
	}
}

// /thread/{slug_or_id}/vote
func ChangeVoteForBranch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	param := params["slug_or_id"]
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		dicts.JsonResponse(w, 500, []byte(err.Error()))
		return
	}
	vote := &models.Vote{}
	err = vote.UnmarshalJSON(body)

	result, err := database.DataManager.MakeThreadVoteDB(vote, param)

	switch err {
	case nil:
		resp, _ := result.MarshalJSON()
		dicts.JsonResponse(w, 200, resp)
	case database.ForumNotFound:
		dicts.JsonResponse(w, 404, dicts.ErrorThread(param))
	case database.UserNotFound:
		dicts.JsonResponse(w, 404, dicts.ErrorFindUserByNick(param))
	default:
		dicts.JsonResponse(w, 500, err.Error())
	}
}
