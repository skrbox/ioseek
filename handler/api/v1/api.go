package v1

import (
	"github.com/gin-gonic/gin"
)

func Registry(r *gin.RouterGroup) {
	r.GET("/session", handleSession)
}
