package main

import (
	_ "GoFrame/boot"
	"GoFrame/cmd"
	"GoFrame/interfaces/http"
)

func main() {
	app := cmd.InitApp() // 构建项目服务

	http.InitHttp(app) // start http

	select {}
}
