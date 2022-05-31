package task

import (
	"time"

	cron "github.com/robfig/cron/v3"

	"github.com/skrbox/ioseek/pkg/conf"
	. "github.com/skrbox/ioseek/pkg/log"
)

var T = cron.New(cron.WithLocation(time.Local), cron.WithLogger(CL))

func init() {
	go syncNew(time.Minute * time.Duration(*conf.TaskSyncNewInterval))
	_, _ = T.AddFunc(conf.Spec[*conf.TaskSyncFullInterval], syncFull)
	_, _ = T.AddFunc(conf.RecommendMonthly, recommend)
}
