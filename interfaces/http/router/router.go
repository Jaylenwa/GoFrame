package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"GoFrame/cmd"
	"GoFrame/interfaces/http/middleware"
	router "GoFrame/interfaces/http/router/conncfg"
)

func Routers(app *cmd.App) *gin.Engine {
	Router := gin.New()

	// 跳过 健康检查
	Router.Use(
		gin.LoggerWithConfig(
			gin.LoggerConfig{
				SkipPaths: []string{"/health/ready", "/health/alive"},
			},
		),
	)
	// 健康检查
	Router.GET("/health/ready", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.String(http.StatusOK, "ready")
	})
	Router.GET("/health/alive", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.String(http.StatusOK, "alive")
	})

	// 配置跨域
	Router.Use(middleware.Cors())
	// 全局recover
	Router.Use(middleware.Recover())
	// 全局错误
	Router.Use(middleware.ErrorHandler)

	ApiGroup := Router.Group("/api/v1")
	router.InitConncfgRouter(ApiGroup, app) // conncfg 模块

	return Router
}
