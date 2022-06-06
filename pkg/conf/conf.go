package conf

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"runtime"
	"strings"

	"github.com/fsnotify/fsnotify"
	. "github.com/skrbox/ioseek/pkg/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)
	if *config != "" {
		viper.SetConfigFile(*config)
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			L.Infof("配置文件发生变更: %s", e.String())
		})
		if err := viper.ReadInConfig(); err != nil {
			L.Error(err)
			os.Exit(1)
		}
	}
	if *showVersion {
		fmt.Println(version())
		os.Exit(0)
	}
}

// 配置信息集合, 集中式管理配置是为了某些配置可能会被循环导入的情况
var (
	// meta
	showVersion   = pflag.BoolP(MetaVersion, "v", false, "查看软件版本")
	config        = pflag.StringP(MetaConfigFile, "c", "", "配置文件")
	_             = pflag.String(MetaListenAddr, ":80", "监听地址")
	_             = pflag.String(MetaUrlPrefix, "/", "统一路由前缀")
	_             = pflag.String(MetaAppName, "修齐方法论", "应用名称")
	metaPlatform  = fmt.Sprintf("%s/%s %s", runtime.GOOS, runtime.GOARCH, runtime.Version())
	metaCommitId  string
	metaVersion   string
	metaBuildAt   string
	metaBranch    string
	metaPoweredBy = `https://github.com/skrbox/ioseek`
	versionTmpl   = `
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
	_ = pflag.String(DBHostPort, "127.0.0.1:3306", "数据库地址:端口")
	_ = pflag.String(DBUserPass, "ioseek:ioseek", "数据库用户:密码")
	_ = pflag.String(DBDatabase, "ioseek", "数据库名")

	// redis
	_ = pflag.String(RedisHostPort, "127.0.0.1:6379", "redis主机:端口")
	_ = pflag.String(RedisPassword, "", "redis密码")

	// task
	_ = pflag.Int64(TaskSyncNewInterval, 30, "同步最新文章周期(分钟),不得低于10")
	_ = pflag.String(TaskSyncFullInterval, Weekly, "全量同步周期(daily|weekly|monthly)")

	// storage
	_ = pflag.String(StorageMediaDir, "", "媒体文件本地目录(如果为空则使用对象存储)")
	_ = pflag.String(StorageAccessKey, "", "对象存储AK")
	_ = pflag.String(StorageSecretKey, "", "对象存储SK")
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
