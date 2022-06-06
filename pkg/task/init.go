package task

import (
	"time"

	cron "github.com/robfig/cron/v3"
	"github.com/spf13/viper"

	c "github.com/skrbox/ioseek/pkg/conf"
	. "github.com/skrbox/ioseek/pkg/log"
)

var T = cron.New(cron.WithLocation(time.Local), cron.WithLogger(CL))

func init() {
	go syncNew(time.Minute * time.Duration(viper.GetInt64(c.TaskSyncNewInterval)))
	_, _ = T.AddFunc(c.Spec[viper.GetString(c.TaskSyncFullInterval)], syncFull)
	_, _ = T.AddFunc(c.RecommendMonthly, recommend)
}
