package main

import (
	"forum/src/database"
	"forum/src/initDB"
	"forum/src/router"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	db := initDB.Init()
	err := db.DbConnect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	database.CreateDataManagerInstance(db.GetConnPool())
	r := router.CreateRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:5000",
		WriteTimeout: 150 * time.Second,
		ReadTimeout:  150 * time.Second,
	}
	println("Listen and Serve " + ":5000")
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
