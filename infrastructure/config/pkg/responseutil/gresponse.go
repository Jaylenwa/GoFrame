package responseutil

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"GoFrame/global"
)

type rspCreateData struct {
	Id interface{} `json:"id"`
}

type rspErrorData struct {
	Cause   interface{} `json:"cause"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

type rspListNull struct {
	Entries    []int `json:"entries"`
	TotalCount int   `json:"total_count"`
}

type rspListExt struct {
	Entries interface{} `json:"entries"`
}

type rspListNullNoTotal struct {
	Entries []int `json:"entries"`
}

// RspOk 返回操作成功
func RspOk(c *gin.Context, code int, any interface{}) {
	switch code {
	case 200:
		c.JSON(
			http.StatusOK,
			any,
		)
	case 201:
		c.JSON(
			http.StatusCreated,
			rspCreateData{
				Id: any,
			},
		)
	case 204:
		c.AbortWithStatus(
			http.StatusNoContent,
		)
	case 200200:
		rspNull := rspListNull{
			Entries:    make([]int, 0),
			TotalCount: 0,
		}
		c.JSON(
			http.StatusOK,
			rspNull,
		)
	case 200201:
		c.JSON(
			http.StatusOK,
			make([]int, 0),
		)
	case 200202:
		rspExt := rspListExt{
			Entries: any,
		}
		c.JSON(
			http.StatusOK,
			rspExt,
		)
	case 200203:
		rspNull := rspListNullNoTotal{
			Entries: make([]int, 0),
		}
		c.JSON(
			http.StatusOK,
			rspNull,
		)

	default:
		c.JSON(
			http.StatusOK,
			any,
		)
	}
}

// RspErr 返回操作失败
func RspErr(c *gin.Context, err error) {
	code := gerror.Code(err)
	if code == gcode.CodeNil && err != nil {
		code = gcode.CodeInternalError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = gerror.NewCode(CommNotFound, err.Error())
		} else {
			err = gerror.NewCode(CommInternalServer, err.Error())
		}
		code = gerror.Code(err)
	}
	subCode := code.(BizCode).BizDetail().SubCode
	msg := msgOld(code.Code())
	rspError(c, code.Code(), subCode, msg, err)
}

func rspError(c *gin.Context, code int, subCode int, message string, err error) {
	c.JSON(
		code,
		rspErrorData{
			//Stack:   gerror.Stack(err), // 堆栈错误
			Cause:   err.Error(),
			Code:    subCode,
			Message: message,
		},
	)
}

var (
	// Languages 支持的语言
	Languages = [3]string{"zh-CN", "zh-TW", "en"}
)

var (
	errorI18n = map[int]map[string]string{
		400: {
			Languages[0]: "参数不合法。",
			Languages[1]: "參數不合法。",
			Languages[2]: "Invalid parameter.",
		},
		401: {
			Languages[0]: "授权无效",
			Languages[1]: "授權無效",
			Languages[2]: "Unauthorized",
		},
		403: {
			Languages[0]: "禁止访问",
			Languages[1]: "禁止訪問",
			Languages[2]: "Forbidden",
		},
		404: {
			Languages[0]: "资源不存在",
			Languages[1]: "資源不存在",
			Languages[2]: "Record does not exist",
		},
		409: {
			Languages[0]: "请求冲突",
			Languages[1]: "請求沖突",
			Languages[2]: "The request is conflicts",
		},
		417: {
			Languages[0]: "请求状态校验失败",
			Languages[1]: "請求狀態校驗失敗",
			Languages[2]: "Failed to request status check",
		},
		500: {
			Languages[0]: "内部错误",
			Languages[1]: "內部錯誤",
			Languages[2]: "Internal Server Error",
		},
	}
)

func msgOld(code int) (msg string) {
	lang := global.Gconfig.Server.Lang
	return errorI18n[code][lang]
}
