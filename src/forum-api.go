package main

import (
	"forum/src/config"
	"forum/src/database"
	"forum/src/initDB"
	"forum/src/router"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	conf := config.GetInstance()
	httpAddr := ":" + strconv.Itoa(conf.Port)

	db := initDB.Init()
	err := db.DbConnect(conf.DbHost, conf.DbDatabase, conf.DbUser, conf.DbPassword, conf.DbPort)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	database.CreateDataManagerInstance(db.GetConnPool())
	r := router.CreateRouter()
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1" + httpAddr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	println("Listen and Serve " + httpAddr)
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
