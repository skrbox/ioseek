package v1

import (
	"github.com/gin-gonic/gin"
)

func Registry(r *gin.RouterGroup) {
	r.POST("/session", handleSession)
	r.GET("/target", handleTarget)
	r.POST("/alert", handleAlert)
}
