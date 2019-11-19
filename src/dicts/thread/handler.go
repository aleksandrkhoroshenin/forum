package thread

import (
	"errors"
	"forum/src/dicts"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateThread(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug_or_id"]
	if slug == "" {
		dicts.MakeResponse(w, 400, errors.New("slug_or_id is empty! "))
		return
	}
}

func CreateThreadBranch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug_or_id"]
	if slug == "" {
		dicts.MakeResponse(w, 400, errors.New("slug_or_id is empty! "))
		return
	}
}

func ChangeBranchDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug_or_id"]
	if slug == "" {
		dicts.MakeResponse(w, 400, errors.New("slug_or_id is empty! "))
		return
	}
}

func GetPostFromBranch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug_or_id"]
	if slug == "" {
		dicts.MakeResponse(w, 400, errors.New("slug_or_id is empty! "))
		return
	}
}

func ChangeVoteForBranch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug_or_id"]
	if slug == "" {
		dicts.MakeResponse(w, 400, errors.New("slug_or_id is empty! "))
		return
	}
}
