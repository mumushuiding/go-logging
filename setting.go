package main

import (
	"encoding/json"
	"log"
	"os"
)

// Configuration 数据库配置结构
type Configuration struct {
	Port            int
	ReadTimeout     int64
	WriteTimeout    int64
	DbType          string `json:"DB_TYPE"`
	MysqlDbHost     string `json:"MYSQL_DB_HOST"`
	MysqlDbPort     string `json:"MYSQL_DB_PORT"`
	MysqlDbUser     string `json:"MYSQL_DB_USER"`
	MysqlDbPassword string `json:"MYSQL_DB_PASSWORD"`
}

// Config 数据库配置
var Config = &Configuration{}

func init() {
	loadConfig()
}
func loadConfig() {
	Config = Config.getConf()
	// 打印配置信息
	config, _ := json.Marshal(&Config)
	log.Printf("configuration:%s", string(config))
}
func (c *Configuration) getConf() *Configuration {
	file, err := os.Open("config.json")
	if err != nil {
		log.Printf("cannot open file config.json：%v", err)
		panic(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(c)
	if err != nil {
		log.Printf("decode config.json failed:%v", err)
		panic(err)
	}
	return c
}
