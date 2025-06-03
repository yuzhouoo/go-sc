package main

import (
	"fmt"
	"go-session-demo/helpers"
	"go-session-demo/routers"
)

func main() {
	// 初始化环境
	_Init()
}

func _Init() {
	// 连接数据库
	DBConnect()

	// web服务
	routers.RouterRun()
}

func DBConnect() {
	db := &helpers.DB{}
	_, err := db.InitDB()

	if err != nil {
		fmt.Println("数据库初始化链接失败")
		return
	}
}
