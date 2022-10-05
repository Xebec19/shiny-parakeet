/*
package accounts contains routes and handlers related to performing actions on accounts
*/
package accounts

import (
	"github.com/Xebec19/shiny-parakeet/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	v1 := r.Group("/v1/account")
	v1.Use(middlewares.Authenticate())
	{
		v1.POST("/create", createAccount)
		v1.GET("/read", readOneAccount)
	}
}
