package middle

import (
	"github.com/axiaoxin-com/logging"
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Registry(e *gin.Engine) {
	e.Use(
		gin.Recovery(),              // panic 恢复
		logging.GinLogger(),         // 请求日志记录
		cors.Default(),              // 跨域处理
		seoProxy(),                  // seo 处理
		analyze(),                   // 统计分析
		authenticate(),              // 管理操作认证
		ginprom.PromMiddleware(nil), // 系统指标暴露
	)
}
