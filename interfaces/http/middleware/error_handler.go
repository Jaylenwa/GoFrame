package middleware

import (
	"GoFrame/global"
	"GoFrame/infrastructure/config/pkg/responseutil"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/pkg/errors"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		// 统一处理 mysql 1062错误
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			responseutil.RspErr(c, gerror.NewCode(responseutil.ComConflictRequest, err.Error()))
			return
		}

		global.GLog.Errorln(fmt.Sprintf("err: %+v", err)) // 记录到日志
		responseutil.RspErr(c, err)                       // 接口错误
	}

}
