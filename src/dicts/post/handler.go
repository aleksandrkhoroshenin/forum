package post

import (
	"forum/src/database"
	"forum/src/dicts"
	"forum/src/dicts/models"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// /post/{id}/details Получение информации о ветке обсуждения
func GetPostDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		dicts.JsonResponse(w, 500, []byte(err.Error()))
		return
	}

	queryParams := r.URL.Query()
	relatedQuery := queryParams.Get("related")
	related := []string{}
	related = append(related, strings.Split(string(relatedQuery), ",")...)

	result, err := database.DataManager.GetPostFullDB(id, related)

	switch err {
	case nil:
		resp, _ := result.MarshalJSON()
		dicts.JsonResponse(w, 200, resp)
	case database.PostNotFound:
		dicts.JsonResponse(w, 404, dicts.ErrorPost(string(id)))
	default:
		dicts.JsonResponse(w, 500, []byte(err.Error()))
	}
}

func ChangePostDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		dicts.JsonResponse(w, 500, []byte(err.Error()))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		dicts.JsonResponse(w, 500, []byte(err.Error()))
		return
	}
	postUpdate := &models.PostUpdate{}
	err = postUpdate.UnmarshalJSON(body)

	if err != nil {
		dicts.JsonResponse(w, 500, []byte(err.Error()))
		return
	}
	result, err := database.DataManager.UpdatePostDB(postUpdate, id)
	switch err {
	case nil:
		resp, _ := result.MarshalJSON()
		dicts.JsonResponse(w, 200, resp)
	case database.PostNotFound:
		dicts.JsonResponse(w, 404, dicts.ErrorPost(string(id)))
	default:
		dicts.JsonResponse(w, 500, []byte(err.Error()))
	}
}
