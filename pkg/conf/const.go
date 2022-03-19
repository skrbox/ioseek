package conf

// 日志输出风格
const (
	Json = "json"
	Txt  = "txt"
)

const (
	Daily   = "daily"
	Weekly  = "weekly"
	Monthly = "monthly"
)

// 任务周期表达式映射
var Spec = map[string]string{
	Daily:   `0 5 * * *`, // 每日凌晨
	Weekly:  `0 4 * * 7`, // 每周一次
	Monthly: `0 3 1 * *`, // 每月一次
}

// 避免打扰到别人，推送时间固定为每月两次
const RecommendMonthly = `0 10 1,16 * *`
