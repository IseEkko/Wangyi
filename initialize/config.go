package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"go.uber.org/zap"
	"singo/conf"
	"singo/global"

	"github.com/spf13/viper"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
	//刚才设置的环境变量 想要生效 我们必须得重启goland
}

func InitConfig() {
	debug := GetEnvInfo("MXSHOP_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("user-web/%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("user-web/%s-debug.yaml", configFilePrefix)
	}
	fmt.Println(configFileName)
	v := viper.New()
	//文件的路径如何设置
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	ServerConfig := conf.ServerConfig{}
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}
	fmt.Println(ServerConfig)
	zap.S().Infof("配置信息：&v", global.ServerConfig)
	fmt.Printf("%v", v.Get("name d"))
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file channed:", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&ServerConfig)
		zap.S().Infof("配置信息：&v", global.ServerConfig)
	})
}
