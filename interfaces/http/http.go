package http

import (
	"GoFrame/cmd"
	"GoFrame/global"
	"GoFrame/infrastructure/config/pkg/util"
	"GoFrame/interfaces/http/router"
	"fmt"
)

// InitHttp Http
func InitHttp(app *cmd.App) {
	// 外部接口
	go func() {
		url := fmt.Sprintf(
			"%s:%s",
			util.ParseHost(global.Gconfig.Server.Host),
			global.Gconfig.Server.PublicPort,
		)
		_ = router.Routers(app).Run(url) // 启动web
	}()
}
