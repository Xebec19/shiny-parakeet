package auth

import (
	"net/http"

	"github.com/Xebec19/shiny-parakeet/util"
	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	var req registerRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	payload := util.ResponseParams{Status: true, Data: util.Stringify(req)}
	c.JSON(http.StatusAccepted, util.Response(payload))
}
