package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

//初始化参数，设置整个项目的全局变量

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

// 从ini文件中获取参数
func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Printf("配置初始化文件出错，'config.ini': %v", err)
	}
	LoadServer(file)
	LoadDataBase(file)
}

// 使用ini库 将设置的参数与变量匹配
// Section会根据[]进行索引，MustString用于设置默认值

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}

func LoadDataBase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("ginVueAdmin")
	DbPassword = file.Section("database").Key("DbPassword").MustString("admin123")
	DbName = file.Section("database").Key("DbName").MustString("ginvueblog")
}
