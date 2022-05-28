package initialize

import (
	"admin/common"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func InitConf(path ...string) {

	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose conf file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv("CONFFILE"); configEnv == "" {
				config = common.ConfigFile
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", config)
			} else {
				config = configEnv
				fmt.Printf("您正在使用configEnv环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用InitConf传递的值,config的路径为%v\n", config) //%v为万能转换
	}
	v := viper.New()
	//设置配置文件路径
	v.SetConfigFile(config)
	//从磁盘中加载配置文件,并且还加载 key/value存储
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error conf file: %s \n", err))
	}
	// 监控配置文件变化并热加载程序
	v.WatchConfig()
	//配置文件发生变更之后会调用的回调函数
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("conf file changed:", e.Name)
		if err = v.Unmarshal(&common.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&common.CONFIG); err != nil {
		fmt.Println(err)
	}
	//赋给全局变量
	common.VP = v
}
