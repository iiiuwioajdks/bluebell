package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func Init() (err error) {
	workdir, _ := os.Getwd()
	viper.SetConfigFile(workdir + "/config.yml")
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
	})
	return err
}
