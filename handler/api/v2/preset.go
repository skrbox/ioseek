package v2

import (
	"github.com/gin-gonic/gin"

	"github.com/skrbox/ioseek/handler"
)

// 预留 v2 版本
func handlePreset(ctx *gin.Context) {
	handler.NewJSONResponse().WithData(gin.H{
		"version":    "2",
		"pathPrefix": "/api/v2/",
	}).Do(ctx)
}
