package main

import (
	"fmt"
	"github.com/phpxin/mdblog/conf"
	"github.com/phpxin/mdblog/tools/logger"
	"github.com/phpxin/mdblog/core"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("usage : ./cm {path of config}")
		os.Exit(1)
	}
	confpath := os.Args[1]

	err := conf.ParseConfig(confpath)
	if err!=nil {
		fmt.Println("err : parse config failed")
		os.Exit(1)
	}

	logger.InitLogger(conf.ConfigInst.Storagepath+"/logs", "20060102")
	core.InitServer()

}
