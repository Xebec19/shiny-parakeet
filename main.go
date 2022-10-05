package main

import (
	"fmt"
	"log"

	"github.com/Xebec19/shiny-parakeet/accounts"
	"github.com/Xebec19/shiny-parakeet/auth"
	db "github.com/Xebec19/shiny-parakeet/db/sqlc"
	"github.com/Xebec19/shiny-parakeet/util"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Connect database
	db.Connect()

	// routes
	auth.Routes(r)
	accounts.Routes(r)

	port := fmt.Sprintf(":%s", config.Port)
	r.Run(port)
}
