package view

import (
	"github.com/gin-gonic/gin"
)

func Registry(r *gin.RouterGroup) {
	r.GET("", handleIndex)
	r.GET("/topic/:id", handleTopic)
	r.GET("/link", handleLink)
	r.GET("/sitemap", handleSitemap)
	r.GET("/sitemap.xml", handleSitemap)
	r.GET("/rss", handleRss)
	r.GET("/rss.html", handleRss)
	r.GET("/feed", handleRss)
	r.GET("/feed.html", handleRss)
}
