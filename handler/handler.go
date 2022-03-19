package handler

import (
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 注册全局接口或视图
// look: 请勿将 `engine` 对象传给用户业务
func Registry(e *gin.Engine, router *gin.RouterGroup) {
	router.GET("/ping", handlePing)
	router.GET("/sys/metrics", ginprom.PromHandler(promhttp.Handler()))
	router.GET("/app/metrics", handleMetrics)
	router.GET("/version", handleVersion)

	// look: 该接口必须在末尾注册
	router.GET("/path", routerPaths(e))
}
