/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"yunfei/cmd"
	"yunfei/internal/bootstrap"
)

func main() {
	// 测试
	//c := cron.New()
	//c.Start()
	//defer c.Stop()
	//
	//_, err := c.AddFunc("@every 3s", func() {
	//	fmt.Println("@every 3s 执行开始")
	//})
	//
	////_, err = c.AddFunc("@every 4s", func() {
	////	fmt.Println("@every 4s 执行开始")
	////})
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//select {}

	// 初始化全局变量 - 未初始化好，会导致事件机制不能正常注册
	bootstrap.Initialization()
	cmd.Execute()
}
