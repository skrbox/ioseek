package conf

import (
	"os"
	"testing"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func TestViperFlag(t *testing.T) {
	os.Setenv("meta.app-name", "anyone")
	viper.AutomaticEnv()
	viper.SetConfigFile("/tmp/ioseek.yml")
	viper.BindPFlags(pflag.CommandLine)
	viper.AutomaticEnv()
	viper.ReadInConfig()
	t.Log(viper.AllSettings())
	t.Log(os.Environ())
	t.Log(viper.Get("test.author"), viper.Get("meta.app-name"))
}
