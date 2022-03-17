package middle

import (
	"github.com/gin-gonic/gin"
)

// 处理各种搜索引擎的 seo 效果
func seoProxy() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// todo: seo 代理
		// todo: seo 上报
		ctx.Next()
	}
}
