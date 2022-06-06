package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/skrbox/ioseek/handler"
	v1 "github.com/skrbox/ioseek/handler/api/v1"
	v2 "github.com/skrbox/ioseek/handler/api/v2"
	"github.com/skrbox/ioseek/handler/middle"
	"github.com/skrbox/ioseek/handler/view"
	_ "github.com/skrbox/ioseek/model"
	c "github.com/skrbox/ioseek/pkg/conf"
	. "github.com/skrbox/ioseek/pkg/log"
	"github.com/skrbox/ioseek/pkg/task"
	"github.com/spf13/viper"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
}

func main() {
	L.Infof("系统应用启动: %s", viper.GetString(c.MetaListenAddr))
	task.T.Start()
	defer task.T.Stop()
	router := gin.New()
	router.LoadHTMLGlob("ui/*.html")
	{
		middle.Registry(router)
		router.HandleMethodNotAllowed = true
		router.NoRoute(handler.Handle404)
		router.NoMethod(handler.Handle405)
		v1.Registry(router.Group(handler.U("/api/v1")))
		v2.Registry(router.Group(handler.U("/api/v2")))
		root := router.Group(handler.U("/"))
		view.Registry(root)

		// 必须最后加载, 且router对象不可传递给业务视图
		handler.Registry(router, root)
	}
	if err := router.Run(viper.GetString(c.MetaListenAddr)); err != nil {
		L.Error(err)
		os.Exit(1)
	}
}
