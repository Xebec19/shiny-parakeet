package accounts

import (
	"net/http"

	"github.com/Xebec19/shiny-parakeet/util"
	"github.com/gin-gonic/gin"
)

func createAccount(c *gin.Context) {
	var req createAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
}
