package user

import (
	"forum/src/database"
	"forum/src/dicts"
	"forum/src/dicts/models"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

// /user/{nickname}/create
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		dicts.JsonResponse(w, 500, err.Error())
		return
	}
	params := mux.Vars(r)
	nickname := params["nickname"]

	user := &models.User{}
	err = user.UnmarshalJSON(body)
	if err != nil {
		dicts.JsonResponse(w, 500, err.Error())
		return
	}
	user.Nickname = nickname
	result, err := database.DataManager.CreateUserDB(user)

	switch err {
	case nil:
		dicts.JsonResponse(w, 201, user)
	case database.UserIsExist:
		dicts.JsonResponse(w, 409, result)
	default:
		dicts.JsonResponse(w, 500, err.Error())
	}
}

// /user/{nickname}/profile
func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	nickname := params["nickname"]

	user, err := database.DataManager.GetUserDB(nickname)
	switch err {
	case nil:
		dicts.JsonResponse(w, 200, user)
	case database.UserNotFound:
		dicts.JsonResponse(w, 404, dicts.ErrorFindUserByNick(nickname))
	default:
		dicts.JsonResponse(w, 500, err.Error())
	}
}

// /user/{nickname}/profile
func ChangeUserInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	nickname := params["nickname"]

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		dicts.JsonResponse(w, 500, err.Error())
		return
	}

	user := &models.User{}
	err = user.UnmarshalJSON(body)
	if err != nil {
		dicts.JsonResponse(w, 500, err.Error())
		return
	}
	user.Nickname = nickname

	err = database.DataManager.UpdateUserDB(user)
	switch err {
	case nil:
		dicts.JsonResponse(w, 200, user)
	case database.UserNotFound:
		dicts.JsonResponse(w, 404, dicts.ErrorFindUserByNick(nickname))
	case database.UserUpdateConflict:
		dicts.JsonResponse(w, 409, dicts.ErrorEmailIsAlreadyExist(nickname))
	default:
		dicts.JsonResponse(w, 500, err.Error())
	}
}
