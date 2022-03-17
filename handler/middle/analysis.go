package middle

import (
	"github.com/gin-gonic/gin"
)

// 统计访问信息
func analyze() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// todo: 统计访问信息等
		ctx.Next()
	}
}
