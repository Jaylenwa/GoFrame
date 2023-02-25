package boot

import (
	"GoFrame/global"
	"GoFrame/infrastructure/config"
	"GoFrame/infrastructure/config/pkg/database/mysql"
	"GoFrame/infrastructure/config/pkg/log"
	_ "GoFrame/infrastructure/config/pkg/validate"
)

// 初始化
func init() {
	global.Gconfig = config.NewConfig() // 初始化全局配置
	global.GLog = log.NewLogger()       // 初始全局化日志
	global.GDB = mysql.NewDB()          // 初始化全局mysql
}
