package main

import (
	"fmt"
	"log"

	"github.com/Xebec19/shiny-parakeet/db"
	"github.com/Xebec19/shiny-parakeet/util"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	db.Connect()
	port := fmt.Sprintf(":%s", config.Port)
	r.Run(port)
}
