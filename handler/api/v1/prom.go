package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/skrbox/ioseek/handler"
)

// 作为 http_sd_config 暴露给 prometheus 建立监控任务
// https://prometheus.io/docs/prometheus/latest/http_sd/
func handleTargetLink(ctx *gin.Context) {
	// todo: ...
	handler.NewJSONResponse().WithData(make([]string, 0)).Do(ctx)
}

// alertmanager 告警全局 webhook, 接收所有告警信息
// 通过annotation中的 send_to 角色来判定发送至谁
// https://prometheus.io/docs/alerting/latest/configuration/#webhook_config
func handleAlert(ctx *gin.Context) {
	// todo: ...
	handler.NewJSONResponse().WithError(handler.NotCompleted).Do(ctx)
}
