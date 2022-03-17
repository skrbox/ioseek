package model

import (
	"fmt"
	"sync"
	"time"

	"github.com/axiaoxin-com/logging"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	c "github.com/skrbox/ioseek/pkg/conf"
	. "github.com/skrbox/ioseek/pkg/log"
)

var (
	DB        *gorm.DB
	once      sync.Once
	mysqlTmpl = `%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`
)

func init() {
	once.Do(func() {
		var driver gorm.Dialector
		if *c.DBHostPort != "" {
			L.Infof("初始化数据连接: %s", *c.DBHostPort)
			driver = mysql.Open(fmt.Sprintf(mysqlTmpl, *c.DBUsername, *c.DBPassword, *c.DBHostPort, *c.DBDatabase))
		} else {
			L.Infof("初始化本地数据库: %s", *c.DBDatabase)
			driver = sqlite.Open(*c.DBDatabase)
		}
		db, err := gorm.Open(driver, &gorm.Config{
			Logger:          logging.NewGormLogger(zap.DebugLevel, zap.DebugLevel, time.Millisecond*500),
			CreateBatchSize: 100,
			NowFunc: func() time.Time {
				return time.Now().Local()
			},
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
				NoLowerCase:   true,
			},
		})
		if err != nil {
			panic(err)
		}
		DB = db
	})
}

type meta struct {
	gorm.Model
	UUID string `gorm:"primaryKey"`
}
