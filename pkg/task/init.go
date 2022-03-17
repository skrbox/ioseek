package task

import (
	"time"

	c "github.com/skrbox/ioseek/pkg/conf"
)

func init() {
	// 定时同步文章信息
	go syncLoop(time.Minute * time.Duration(*c.TaskSyncInterval))
}
