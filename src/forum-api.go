package main

import (
	"forum/src/config"
	"forum/src/database"
	"forum/src/initDB"
	"forum/src/router"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"strconv"
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
	println("Listen and Serve " + httpAddr)
	if err := fasthttp.ListenAndServe(httpAddr, r.Handler); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
