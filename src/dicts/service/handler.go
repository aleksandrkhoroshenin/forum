package service

import (
	"forum/src/database"
	"forum/src/dicts"
	"net/http"
)

func ClearDB(w http.ResponseWriter, r *http.Request) {
	database.DataManager.ClearDB()
}

func GetInformationDB(w http.ResponseWriter, r *http.Request) {
	result, err := database.DataManager.GetStatusDB()

	switch err {
	case nil:
		dicts.JsonResponse(w, 200, result)
	default:
		dicts.JsonResponse(w, 500, err.Error())
	}
}
