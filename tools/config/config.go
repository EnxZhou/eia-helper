package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var AmapKey string
var cfgSearch *viper.Viper

//载入配置文件
func ConfigSetup(path string) {
	viper.SetConfigFile(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}

	//Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}
	// 启动参数
	AmapKey = viper.GetString("key")
	if AmapKey == "" {
		panic("need valid amap key")
	}

	cfgSearch = viper.Sub("settings.search")
	if cfgSearch == nil {
		panic("config not found settings.search")
	}
	SearchConfig = InitSearch(cfgSearch)
}

func SetConfig(configPath string, key string, value interface{}) {
	viper.AddConfigPath(configPath)
	viper.Set(key, value)
	_ = viper.WriteConfig()
}
