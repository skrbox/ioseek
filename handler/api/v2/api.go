package v2

import (
	"github.com/gin-gonic/gin"
)

func Registry(r *gin.RouterGroup) {
	r.GET("", handlePreset)
}
