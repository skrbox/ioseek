package conf

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

// 配置参数集
const (
	MetaVersion    = "version"
	MetaConfigFile = "meta.config-file"
	MetaListenAddr = "meta.listen-addr"
	MetaUrlPrefix  = "meta.url-prefix"
	MetaAppName    = "meta.app-name"

	DBHostPort = "db.host-port"
	DBUserPass = "db.user-pass"
	DBDatabase = "db.database"

	TaskSyncNewInterval  = "task.sync-new-interval"
	TaskSyncFullInterval = "task.sync-full-interval"

	RedisHostPort = "redis.host-port"
	RedisPassword = "redis.password"

	StorageMediaDir  = "storage.media-dir" // 如果设置了本地目录则不使用对象存储
	StorageAccessKey = "storage.access-key"
	StorageSecretKey = "storage.secret-key"
)
