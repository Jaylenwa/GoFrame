package assembler

import (
	dto "GoFrame/application/dto/conncfg"
	"GoFrame/domain/entity/conncfg"
)

// ConncfgRsp 请求参数
type ConncfgRsp struct {
}

// NewConncfgRsp 实例化
func NewConncfgRsp() *ConncfgRsp {
	return &ConncfgRsp{}
}

// E2DCreateConncfg entity转换成dto
func (a *ConncfgRsp) E2DCreateConncfg(en *entity.Conncfg) *dto.CreateConncfgRsp {
	var rspDto dto.CreateConncfgRsp
	rspDto.Id = en.ConncfgId

	return &rspDto
}
