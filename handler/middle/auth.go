package middle

import (
	"github.com/gin-gonic/gin"
)

// 管理员身份统一认证，视图层无需处理
func authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// todo: 对管理操作进行身份认证
		ctx.Next()
	}
}
