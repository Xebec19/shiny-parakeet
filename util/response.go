package util

import (
	"log"

	"github.com/gin-gonic/gin"
)

type ResponseParams struct {
	Status bool
	Data   interface{}
}

func ErrorResponse(err error) gin.H {
	log.Printf("%s", err.Error())
	return gin.H{"error": err.Error()}
}

func Response(payload ResponseParams) gin.H {
	return gin.H{"status": payload.Status, "data": payload.Data}
}
