package model

import (
	"github.com/phpxin/mdblog/conf"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func InitModel() {
	//"user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	connstr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", conf.ConfigInst.Dbuser,
		conf.ConfigInst.Dbpassword,
		conf.ConfigInst.Dbaddr,
		conf.ConfigInst.Dbname)

	var err error
	db, err = gorm.Open("mysql", connstr)
	if err!=nil {
		panic(err)
	}

}
