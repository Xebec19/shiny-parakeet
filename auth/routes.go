package auth

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	v1 := r.Group("/v1/auth")
	{
		v1.POST("/register", register)
	}
}
