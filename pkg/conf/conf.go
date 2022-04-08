package conf

import (
	"bytes"
	"fmt"
	"html/template"
	"runtime"
	"strings"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func init() {
	kingpin.HelpFlag.Short('h')
	kingpin.Version(version()).VersionFlag.Short('v')
	kingpin.Parse()
}

// 配置信息集合, 集中式管理配置是为了某些配置可能会被循环导入的情况
var (
	// meta
	MetaListenAddr = kingpin.Flag("meta.listen-addr", "监听地址").Default(":80").String()
	MetaUrlPrefix  = kingpin.Flag("meta.url-prefix", "统一路由前缀").Default("/").String()
	MetaAppName    = kingpin.Flag("meta.app-name", "应用名称").Default("修齐方法论").String()
	metaPlatform   = fmt.Sprintf("%s/%s %s", runtime.GOOS, runtime.GOARCH, runtime.Version())
	metaCommitId   string
	metaVersion    string
	metaBuildAt    string
	metaBranch     string
	metaPoweredBy  = `https://github.com/skrbox/ioseek`
	versionTmpl    = `
powered by {{.poweredBy}}
    version:	{{.version}}
    branch:	{{.branch}}
    revision:	{{.commitId}}
    buildAt:	{{.buildAt}}
    platform:	{{.platform}}
`
	MetaVersionMap = map[string]string{
		"version":   metaVersion,
		"commitId":  metaCommitId,
		"branch":    metaBranch,
		"buildAt":   metaBuildAt,
		"platform":  metaPlatform,
		"poweredBy": metaPoweredBy,
	}

	// db
	DBUsername = kingpin.Flag("db.username", "数据库用户名").Default("ioseek").String()
	DBPassword = kingpin.Flag("db.password", "数据库密码").Default("ioseek.cn").String()
	DBHostPort = kingpin.Flag("db.host-port", "数据库连接地址，如[127.0.0.1:3306], 当为空时使用sqlite").String()
	DBDatabase = kingpin.Flag("db.database", "数据库名称, 当hostProt为空则作为sqlite数据库文件").Default("ioseek.db").String()

	// task
	TaskSyncNewInterval  = kingpin.Flag("task.sync-new-interval", "同步最新文章周期(分钟)").Default("60").Int64()
	TaskSyncFullInterval = kingpin.Flag("task.sync-full-interval", "全量同步周期").Default(Weekly).Enum(Daily, Weekly, Monthly)
)

// 构建命令行版本输出信息
func version() string {
	t := template.Must(template.New("version").Parse(versionTmpl))

	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, "version", MetaVersionMap); err != nil {
		panic(err)
	}
	return strings.TrimSpace(buf.String())
}
