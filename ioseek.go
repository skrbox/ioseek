package main

import (
	"fmt"
	"runtime"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	_ "github.com/skrbox/ioseek/task"
)

func init() {

}

var (
	// 项目元信息
	webpage  = "https://ioseek.cn/"
	app      = "ioseek"
	commitId string
	version  string
	buildAt  string
	branch   string
	platform = fmt.Sprintf("%s/%s %s", runtime.GOOS, runtime.GOARCH, runtime.Version())

	// 项目启动时配置
	webListenAddr  = kingpin.Flag("meta.listen-addr", "监听地址").Default(":80").String()
	metaConfigFile = kingpin.Flag("meta.config-file", "配置文件路径").Default("ioseek.yml").String()
)

func main() {

}
