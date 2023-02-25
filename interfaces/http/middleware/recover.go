package middleware

import (
	"GoFrame/infrastructure/config/pkg/responseutil"
	"GoFrame/infrastructure/config/pkg/util"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"

	"GoFrame/global"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				errorInfo := responseutil.Panic(fmt.Sprintf("%s", err))
				global.GLog.Errorln(util.PrintJson(errorInfo.Internale))                                             // 记录到日志
				responseutil.RspErr(c, gerror.NewCode(responseutil.CommInternalServer, gconv.String(errorInfo.Out))) // 前端返回
				return
			}
		}()
		c.Next()
	}
}
