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

func CreateForum(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		dicts.MakeResponse(w, 500, err.Error())
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
		dicts.MakeResponse(w, 201, forum)
	case database.UserNotFound:
		dicts.MakeResponse(w, 404, dicts.MakeErrorUser(forum.User))
	case database.ForumIsExist:
		dicts.MakeResponse(w, 409, forum)
	default:
		dicts.MakeResponse(w, 500, err.Error())
	}
}

func GetBranchDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug"]
	forum, err := database.DataManager.GetForumDB(slug)

	switch err {
	case nil:
		dicts.MakeResponse(w, 200, forum)
	case database.ForumNotFound:
		dicts.MakeResponse(w, 404, dicts.MakeErrorForum(slug))
	default:
		dicts.MakeResponse(w, 500, err.Error())
	}
}

func CreateForumBranch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug"]
	if slug == "" {
		dicts.MakeResponse(w, 400, errors.New("slug is empty! "))
		return
	}
	//database.DataManager.CreateThreadDB()
}

func GetBranchThreads(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug"]
	if slug == "" {
		dicts.MakeResponse(w, 400, errors.New("slug is empty! "))
		return
	}
}

func GetBranchUsers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug"]
	if slug == "" {
		dicts.MakeResponse(w, 400, errors.New("slug is empty! "))
		return
	}
}
