package main

import (
	"forum/src/config"
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

	log.Println(conf)
	r := router.CreateRouter()

	if err := fasthttp.ListenAndServe(httpAddr, r.Handler); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
