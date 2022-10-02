package auth

import (
	"database/sql"
	"net/http"

	db "github.com/Xebec19/shiny-parakeet/db/sqlc"
	"github.com/Xebec19/shiny-parakeet/util"
	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	var req registerRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	args := db.CreateUserParams{
		FirstName: req.FirstName,
		LastName:  sql.NullString{String: req.LastName, Valid: true},
		Email:     req.Email,
		Password:  req.Password,
	}
	user, err := util.DBQuery.CreateUser(c, args)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	payload := util.ResponseParams{
		Status: true,
		Data:   user,
	}
	c.JSON(http.StatusAccepted, util.Response(payload))
}
