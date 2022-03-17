package conf

import (
	"fmt"
	"runtime"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func init() {
	kingpin.Parse()
}

// 配置信息集合, 集中式管理配置是为了某些配置可能会被循环导入的情况
var (
	// meta
	MetaListenAddr = kingpin.Flag("meta.listen-addr", "监听地址").Default(":80").String()
	MetaUrlPrefix  = kingpin.Flag("meta.url-prefix", "统一路由前缀").Default("/").String()
	MetaPlatform   = fmt.Sprintf("%s/%s %s", runtime.GOOS, runtime.GOARCH, runtime.Version())
	MetaWebPage    = kingpin.Flag("meta.web-page", "应用web地址").Default("https://ioseek.cn").String()
	MetaAppName    = kingpin.Flag("meta.app-name", "应用名称").Default("ioseek").String()
	MetaCommitId   string
	MetaVersion    string
	MetaBuildAt    string
	MetaBranch     string

	// db
	DBUsername = kingpin.Flag("db.username", "数据库用户名").Default("ioseek").String()
	DBPassword = kingpin.Flag("db.password", "数据库密码").Default("ioseek.cn").String()
	DBHostPort = kingpin.Flag("db.host-port", "数据库连接地址，如[127.0.0.1:3306], 当为空时使用sqlite").String()
	DBDatabase = kingpin.Flag("db.database", "数据库名称, 当hostProt为空则作为sqlite数据库文件").Default("ioseek.db").String()

	// log
	LogStyle = kingpin.Flag("log.style", "日志风格").Default(Json).Enum(Json, Txt)

	// task
	TaskSyncInterval = kingpin.Flag("task.sync-interval", "同步文章周期(分钟)").Default("60").Int64()
)
