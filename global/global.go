package global

import (
	"GoFrame/infrastructure/config"
	"GoFrame/infrastructure/config/pkg/log"
	"gorm.io/gorm"
)

var (
	Gconfig *config.Config // 全局配置
	GLog    log.Logger     // 全局log
	GDB     *gorm.DB       // 全局 DB-edgp
)
