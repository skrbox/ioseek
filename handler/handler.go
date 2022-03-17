package handler

import (
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 注册全局接口或视图
func Registry(router *gin.RouterGroup) {
	router.GET("/ping", handlePing)
	router.GET("/sys/metrics", ginprom.PromHandler(promhttp.Handler()))
	router.GET("/app/metrics", handleMetrics)
	router.GET("/version", handleVersion)
}
