package auth

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	db "github.com/Xebec19/shiny-parakeet/db/sqlc"
	"github.com/Xebec19/shiny-parakeet/token"
	"github.com/Xebec19/shiny-parakeet/util"
	"github.com/gin-gonic/gin"
)

// register creates a new user in db and returns her user id
func register(c *gin.Context) {
	var req registerRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	hashPassword, err := util.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	args := db.CreateUserParams{
		FirstName: req.FirstName,
		LastName:  sql.NullString{String: req.LastName, Valid: true},
		Email:     req.Email,
		Password:  hashPassword,
	}
	user, err := db.DBQuery.CreateUser(c, args)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	payload := util.ResponseParams{
		Status: true,
		Data:   user.UserID,
	}
	c.JSON(http.StatusAccepted, util.Response(payload))
}

// login takes email and password, and generates a jwt token if credentials are correct
func login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	user, err := db.DBQuery.FindUser(c, req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	// check password
	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	// assign token
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	maker, err := token.NewJWTMaker(config.JWTSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	token, err := maker.CreateToken(user.UserID.String(), 24*time.Hour)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	responsePayload := util.ResponseParams{
		Status: true,
		Data:   token,
	}
	c.JSON(http.StatusAccepted, util.Response(responsePayload))
}
