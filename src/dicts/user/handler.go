package user

import (
	"errors"
	"forum/src/database"
	"forum/src/dicts"
	"forum/src/dicts/models"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		dicts.MakeResponse(w, 500, err.Error())
		return
	}
	params := mux.Vars(r)
	nickname := params["nickname"]
	if nickname == "" {
		dicts.MakeResponse(w, 400, errors.New("nickname is empty! "))
		return
	}
	user := &models.User{}
	err = user.UnmarshalJSON(body)
	if err != nil {
		return
	}
	user.Nickname = nickname
	if users, err := database.DataManager.CreateUserDB(user); err != nil {
		log.Println(users, err)
		return
	}
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	nickname := params["nickname"]
	if nickname == "" {
		dicts.MakeResponse(w, 400, errors.New("nickname is empty! "))
		return
	}
	user, err := database.DataManager.GetUserDB(nickname)
	switch err {
	case nil:
		dicts.MakeResponse(w, 200, user)
	case database.UserNotFound:
		dicts.MakeResponse(w, 404, dicts.MakeErrorUser(nickname))
	default:
		dicts.MakeResponse(w, 500, err.Error())
	}
}

func ChangeUserInfo(w http.ResponseWriter, r *http.Request) {

}
