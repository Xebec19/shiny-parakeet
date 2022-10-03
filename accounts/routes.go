package accounts

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	v1 := r.Group("/v1/account")
	{
		v1.POST("/create", createAccount)
	}
}
