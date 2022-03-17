package view

import (
	"github.com/gin-gonic/gin"

	"github.com/skrbox/ioseek/handler"
)

func handleSitemap(ctx *gin.Context) {
	// todo: 站点地图生成
	handler.NewJsonResponse().WithError(handler.NotCompleted).Do(ctx)
}

func handleRss(ctx *gin.Context) {
	// todo: rss 订阅
	handler.NewJsonResponse().WithError(handler.NotCompleted).Do(ctx)
}
