package handler

import (
	"GoFrame/infrastructure/config/pkg/responseutil"
	"GoFrame/infrastructure/config/pkg/validate"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gerror"

	adapter "GoFrame/application/adapter/conncfg"
	dto "GoFrame/application/dto/conncfg"
)

type ConncfgHandler struct {
	ConncfgSrv *adapter.ConncfgAdapter
}

// CreateConncfg 配置创建
func (h *ConncfgHandler) CreateConncfg(c *gin.Context) {
	// 参数解析
	dtoConncfg := &dto.CreateConncfgReq{}
	err := c.BindJSON(dtoConncfg)
	if err != nil {
		err = gerror.NewCode(responseutil.CodeCfgBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 参数过滤

	err = validate.Validate(dtoConncfg)
	if err != nil {
		err = gerror.NewCode(responseutil.CodeCfgBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 业务处理
	ctx := context.Background()
	ctx = context.WithValue(ctx, "clientIp", c.ClientIP())
	res, err := h.ConncfgSrv.AddConcfg(ctx, dtoConncfg)
	if err != nil {
		_ = c.Error(err)
		return
	}
	responseutil.RspOk(c, http.StatusCreated, res.Id)
}
