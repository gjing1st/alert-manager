package config

import (
	configor "github.com/jinzhu/configor"
)

// Config 存储全局参数，供其他模块使用
var Config = struct {
	Log struct {
		Output string `default:"std"`  //日志输出，标准输出或文件
		Level  string `default:"info"` //日志等级
		Caller bool   `default:"true"` //是否打印调用者信息
		Dir    string `default:"."`    //存放目录
	}
	Web struct {
		Port string `default:"9680"`
		Cors bool   `default:"true"`
	}
	Mysql struct {
		Host     string `default:"127.0.0.1"`
		UserName string `default:"root"`
		Password string `default:"root"`
		DBName   string `default:"publish"`
		Port     string `default:"3306"`
		MinConns int    `default:"90"`  //连接池最小连接数量 不要太小
		MaxConns int    `default:"120"` //连接池最大连接数量 两者相差不要太大
	}
	StorePath   string `default:"./upload/file"`
	CrontabTime int    `default:"60"`
}{}

// InitConfig 读取用户的配置文件
func InitConfig() {
	err := configor.Load(&Config, "./config/config.yml")
	if err != nil {
		panic("config load error" + err.Error())
	}
}
