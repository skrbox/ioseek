package task

import (
	"time"

	. "github.com/skrbox/ioseek/pkg/log"
)

// 定时同步文章信息
func syncNew(interval time.Duration) {
	go func(interval time.Duration) {
		if interval < time.Minute*10 || interval > time.Hour*3 {
			L.Warnf("推荐同步周期为 10m-3h, 当前周期可能不太合理: %s", interval.String())
		}
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case point := <-ticker.C:
				L.Infof("开始同步最新内容")
				// todo: 同步逻辑
				// m.DB.Find()
				L.Infow("一轮同步结束", "duration", time.Since(point).String())
			}
		}
	}(interval)
}

// 全量同步纠正
func syncFull() {
	// todo: ...
}
