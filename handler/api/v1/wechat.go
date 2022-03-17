package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/skrbox/ioseek/handler"
)

// 微信公众号会话控制
func handleSession(ctx *gin.Context) {
	// todo: 会话控制
	handler.NewJsonResponse().WithError(handler.NotCompleted).Do(ctx)
}
