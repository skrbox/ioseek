package main

import (
	"fmt"
	_ "github.com/pmeta/ioseek/task"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"runtime"
)

func init() {

}

// 项目元信息
var (
	webpage  = "https://ioseek.cn/"
	app      = "ioseek"
	commitId string
	version  string
	buildAt  string
	branch   string
	platform = fmt.Sprintf("%s/%s %s", runtime.GOOS, runtime.GOARCH, runtime.Version())

	webListenAddr  = kingpin.Flag("meta.listen-addr", "监听地址").Default(":80").String()
	metaConfigFile = kingpin.Flag("meta.config-file", "配置文件路径").Default("ioseek.yml").String()
)

func main() {

}
