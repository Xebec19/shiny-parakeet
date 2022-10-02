package auth

import (
	"net/http"

	db "github.com/Xebec19/shiny-parakeet/db"
	"github.com/Xebec19/shiny-parakeet/util"
	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	var req registerRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	user := db.Users{FirstName: req.FirstName, LastName: req.LastName, Email: req.Email, Password: req.Password}
	db.DB.Create(&user)
	payload := util.ResponseParams{Status: true, Data: user}
	c.JSON(http.StatusAccepted, util.Response(payload))
}
