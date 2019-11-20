package forum

import (
	"errors"
	"forum/src/database"
	"forum/src/dicts"
	"forum/src/dicts/models"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

// /forum/create
func CreateForum(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		dicts.JsonResponse(w, 500, err.Error())
		return
	}
	forum := &models.Forum{}
	err = forum.UnmarshalJSON(body)
	if err != nil {
		return
	}
	err = database.DataManager.CreateForumDB(forum)
	switch err {
	case nil:
		dicts.JsonResponse(w, 201, forum)
	case database.UserNotFound:
		dicts.JsonResponse(w, 404, dicts.ErrorFindUserByNick(forum.User))
	case database.ForumIsExist:
		dicts.JsonResponse(w, 409, forum)
	default:
		dicts.JsonResponse(w, 500, err.Error())
	}
}

// /forum/{slug}/details
func GetBranchDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug"]
	forum, err := database.DataManager.GetForumDB(slug)

	switch err {
	case nil:
		dicts.JsonResponse(w, 200, forum)
	case database.ForumNotFound:
		dicts.JsonResponse(w, 404, dicts.ErrorFindForumWithSlug(slug))
	default:
		dicts.JsonResponse(w, 500, err.Error())
	}
}

// /forum/{slug}/create
func CreateForumBranch(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		dicts.JsonResponse(w, 500, err.Error())
		return
	}
	params := mux.Vars(r)
	slug := params["slug"]
	if slug == "" {
		dicts.JsonResponse(w, 400, errors.New("slug is empty! "))
		return
	}
	thread := &models.Thread{}
	err = thread.UnmarshalJSON(body)
	if err != nil {
		dicts.JsonResponse(w, 500, err.Error())
		return
	}
	thread.Slug = slug
	err = database.DataManager.CreateThreadDB(thread)
	switch err {
	case nil:
		dicts.JsonResponse(w, 201, thread)
	case database.ForumOrAuthorNotFound:
		dicts.JsonResponse(w, 404, dicts.ErrorFindUserByNick(slug))
	case database.ThreadIsExist:
		dicts.JsonResponse(w, 409, thread)
	default:
		dicts.JsonResponse(w, 500, err.Error())
	}
}

// /forum/{slug}/threads
func GetBranchThreads(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug"]
	if slug == "" {
		dicts.JsonResponse(w, 400, errors.New("slug is empty! "))
		return
	}
	queryParams := r.URL.Query()
	query := dicts.ParseQueryParams(queryParams)

	threads, err := database.DataManager.GetForumThreads(slug, query)
	switch err {
	case nil:
		dicts.JsonResponse(w, 200, threads)
	case database.ForumNotFound:
		dicts.JsonResponse(w, 404, dicts.ErrorFindForumWithSlug(slug))
	default:
		dicts.JsonResponse(w, 500, err.Error())
	}
}

// /forum/{slug}/users
func GetBranchUsers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug"]
	if slug == "" {
		dicts.JsonResponse(w, 400, errors.New("slug is empty! "))
		return
	}
	queryParams := r.URL.Query()
	query := dicts.ParseQueryParams(queryParams)
	users, err := database.DataManager.GetForumUsersDB(slug, query)
	switch err {
	case nil:
		dicts.JsonResponse(w, 200, users)
	case database.ForumNotFound:
		dicts.JsonResponse(w, 404, dicts.ErrorFindForumWithSlug(slug))
	default:
		dicts.JsonResponse(w, 500, err.Error())
	}
}
