package view

import (
	"github.com/gin-gonic/gin"

	"github.com/skrbox/ioseek/handler"
)

func handleIndex(ctx *gin.Context) {
	// todo: 首页
	handler.NewHTMLResponse().WithError(handler.NotCompleted).Do(ctx)
}

func handleTopic(ctx *gin.Context) {
	// todo: 话题列表页
	handler.NewHTMLResponse().WithError(handler.NotCompleted).Do(ctx)
}

func handleLink(ctx *gin.Context) {
	// todo: 友链页
	handler.NewHTMLResponse().WithError(handler.NotCompleted).Do(ctx)
}
