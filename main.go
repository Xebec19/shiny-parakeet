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

// @title         assignment
// @version       1.0
// @description   This is a Go application having JWT authentication, Unit tests,etc using postgresql as database
// @host          localhost:8080
// @BasePath      /
// @schemes       http
// @contact.name  Rohan Kumar Thakur
// @contact.email rohandeveloper106@gmail.com
// @license.name  GNU GENERAL PUBLIC LICENSE
func main() {
	r := gin.Default()
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Connect database
	db.Connect()

	// Routes
	auth.Routes(r)
	accounts.Routes(r)

	port := fmt.Sprintf(":%s", config.Port)
	// Start server
	r.Run(port)
}
