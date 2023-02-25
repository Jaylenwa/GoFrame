package assembler

import (
	dto "GoFrame/application/dto/conncfg"
	"GoFrame/domain/entity/conncfg"
)

// ConncfgReq 请求参数
type ConncfgReq struct {
}

// NewConncfgReq NewUConncfgReq
func NewConncfgReq() *ConncfgReq {
	return &ConncfgReq{}
}

// D2ECreateConncfg dto转换成entity
func (a *ConncfgReq) D2ECreateConncfg(dto *dto.CreateConncfgReq) *entity.Conncfg {
	var rspEn entity.Conncfg

	rspEn.LinkName = dto.LinkName
	rspEn.LinkDesc = dto.LinkDesc
	rspEn.ServerName = dto.ServerName
	rspEn.Client = dto.Client
	rspEn.User = dto.User
	rspEn.Pwd = dto.Pwd
	rspEn.Sid = dto.Sid
	rspEn.Sysnr = dto.Sysnr
	rspEn.CharacterSet = dto.CharacterSet
	rspEn.Language = dto.Language
	rspEn.CfgSource = 1
	rspEn.State = dto.State

	return &rspEn
}
