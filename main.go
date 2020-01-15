package main

import (
	"fmt"
	"os"

	"github.com/phpxin/mdblog/conf"
	"github.com/phpxin/mdblog/controllers"
	"github.com/phpxin/mdblog/core"
	model "github.com/phpxin/mdblog/models"
	"github.com/phpxin/mdblog/tools/logger"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("usage : ./cm {path of config}")
		os.Exit(1)
	}
	confpath := os.Args[1]

	// 1. 解析配置文件
	err := conf.ParseConfig(confpath)
	if err != nil {
		fmt.Println("err : parse config failed", err)
		os.Exit(1)
	}

	// 2. 初始化日志
	logger.InitLogger(conf.ConfigInst.Storagepath+"/logs", "20060102")

	// 3. 模型初始化
	model.InitModel()

	// 4. 初始化文档结构树
	err = core.GenerateTreeFolder()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// 5. 加载路由
	core.Router(&controllers.BlogController{})
	core.Router(&controllers.IndexController{})
	core.Router(&controllers.TestController{})
	core.Router(&controllers.AdminController{})

	// 6. 初始化 WEB 服务
	core.InitServer()

}
