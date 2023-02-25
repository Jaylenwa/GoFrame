package router

import (
	"github.com/gin-gonic/gin"

	"GoFrame/cmd"
	"GoFrame/interfaces/http/handler/conncfg"
)

func InitConncfgRouter(Router *gin.RouterGroup, app *cmd.App) {
	hand := handler.ConncfgHandler{
		ConncfgSrv: app.ConncfgSrv,
	}
	UserRouter := Router.Group("/cfg")
	{
		UserRouter.POST("/conncfg", hand.CreateConncfg) // 创建Conncfg
	}
}
