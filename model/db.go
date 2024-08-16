package model

import (
	"fmt"
	"gin-vue-blog/utils"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var err error

func InitDB() {
	timeout := "10s"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
		timeout,
	)
	mysqllogger := logger.Default.LogMode(logger.Info)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqllogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Printf("mysql连接出错，%v", err)
	}
	_ = db.AutoMigrate(&User{}, &Article{}, &Category{}, &Profile{}, &Comments{}, &FriendLink{})
}
