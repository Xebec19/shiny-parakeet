/*
package middlewares contains middlewares required for authenticating, or retreiving data from a request
*/
package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/Xebec19/shiny-parakeet/token"
	"github.com/Xebec19/shiny-parakeet/util"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		config, err := util.LoadConfig(".")
		if err != nil {
			log.Fatal("cannot load config:", err)
		}

		maker, _ := token.NewJWTMaker(config.JWTSecret)
		tokenHeader := c.Request.Header["Authorization"]
		if len(tokenHeader) == 0 {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		token := strings.Split(tokenHeader[0], " ")[1]
		// verify token
		payload, err := maker.VerifyToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		// save payload
		c.Set("userId", payload.Username)

		c.Next()
	}
}
