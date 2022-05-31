package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/skrbox/ioseek/handler"
)

// 微信公众号会话控制，由中间件处理认证，此处只需实现逻辑
func handleSession(ctx *gin.Context) {
	// todo: 会话控制
	handler.NewJSONResponse().WithError(handler.NotCompleted).Do(ctx)
}

// 友链自助添加，？需要管理员审核内容

// 绑定管理员
// 第一个绑定的管理员为所有者，已有超管无法再次绑定
// 认证口令由所有者申请，系统初始化后返回，客户端输入口令绑定后失效
// 超管可禁用普通管理员
