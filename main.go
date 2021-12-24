package main

import (
	"ApiForCode/bootstrap"
	"ApiForCode/global"
	"ApiForCode/utils"
	"log"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()

	//初始化Ocr识别库
	err := utils.InitDll()
	if err != nil {
		log.Println("初始化Ocr识别失败！")
		return
	}

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")
	// 启动服务器
	bootstrap.RunServer()
}
