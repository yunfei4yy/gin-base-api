package bootstrap

import (
	"fmt"
	"log"
	"yunfei/app/service"
	"yunfei/app/task"
	"yunfei/internal/config"
	"yunfei/internal/crontab"
	"yunfei/internal/event"
	"yunfei/internal/global"
	"yunfei/internal/logger"
	"yunfei/internal/mysql"
	"yunfei/internal/validator"
)

func Initialization() {
	fmt.Println("Initializing application...") // 调试日志
	var err error
	// 初始化配置文件
	global.Config = config.GetConfig()

	// 初始化数据库
	global.DB = mysql.GetConnection()
	if global.DB == nil {
		log.Fatal("数据库连接未初始化")
	}

	// 初始化日志
	if global.Logger, err = logger.New(); err != nil {
		fmt.Println("初始化日志错误: ", err)
		panic(err)
		return
	}

	//// 下面是测试代码用完删除
	//global.Logger.Info("log_init", zap.String("name", "yunfei"))
	//// 下边是测试代码，用完删除
	//for {
	//
	//	global.Logger.Info("info", zap.String("name", "eve"))
	//	global.Logger.Debug("debug", zap.String("name", "eve"))
	//	global.Logger.Error("error", zap.String("name", "eve"))
	//}
	//
	//var adminList model.Admin
	//global.DB.First(&adminList)
	//fmt.Println(adminList)
	service.InitServices() // 因为之前的版本初始化数据库连接之前就已经初始化了service，所以全局变量a.Db是nil，所以在数据库连接初始化之后再更新db这样就不会nil了
	// 注册验证器
	validator.InitValidator()
	// 初始化事件机制
	global.EventDispatcher = event.New()
	// 初始化定时任务
	global.Crontab = crontab.Init()
	global.Crontab.AddTask(task.Task()...)
	global.Crontab.Start()
}
