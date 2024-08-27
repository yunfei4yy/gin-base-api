package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"yunfei/internal/global"
)

func GetConnection() (db *gorm.DB) {
	// 获取数据库配置
	config := global.Config.Database

	// 拼装dsn
	dsn := config.UserName + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Database + "?charset=" + config.Charset + "&parseTime=True&loc=Local"
	// 初始化配置
	mysqlConfig := mysql.Config{
		DSN: dsn,
	}
	var err error
	// 连接数据库
	db, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // 如果是true代表禁用表名复数
		},
	})
	if err != nil {
		log.Println("数据库连接失败: ", err)
		panic(err)
	}

	//global.DB = db
	return db
}
