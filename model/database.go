package model

import (
	"fmt"
	"log"
	"time"

	// mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/mumushuiding/go-logging/config"
)

var db *gorm.DB

// Model 其它数据结构的公共部分
type Model struct {
	ID        int `gorm:"primary_key" json:"id"`
	Createdon int `json:"createdon"`
}

// 配置
var conf = *config.Config

// Setup 初始化一个db连接
func Setup() {
	var err error
	log.Println("数据库初始化！！")
	db, err = gorm.Open(conf.DbType, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName))
	if err != nil {
		log.Fatalf("数据库连接失败 err: %v", err)
	}
	// 启用Logger，显示详细日志
	db.LogMode(true)

	db.SingularTable(true) //全局设置表名不可以为复数形式
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.DB().SetMaxIdleConns(conf.DbMaxIdleConns)
	db.DB().SetMaxOpenConns(conf.DbMaxOpenConns)

	db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(&Logdata{})

}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("Createdon"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}
	}
}
