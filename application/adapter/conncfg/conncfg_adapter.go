package adapter

import (
	assembler "GoFrame/application/assember/aconncfg"
	"GoFrame/domain/service/conncfg"
	"GoFrame/domain/service/conncfg/impl"
	"context"

	dto "GoFrame/application/dto/conncfg"
)

type ConncfgAdapter struct {
	conncfgReq *assembler.ConncfgReq
	conncfgRsp *assembler.ConncfgRsp
	conncfgSvc service.ConncfgService
}

func NewConncfgAdapter() *ConncfgAdapter {
	return &ConncfgAdapter{
		conncfgReq: assembler.NewConncfgReq(),
		conncfgRsp: assembler.NewConncfgRsp(),
		conncfgSvc: impl.NewConncfgServiceImpl(),
	}
}

// AddConcfg 配置添加
func (s *ConncfgAdapter) AddConcfg(ctx context.Context, req *dto.CreateConncfgReq) (*dto.CreateConncfgRsp, error) {
	var rsp dto.CreateConncfgRsp
	encfg := s.conncfgReq.D2ECreateConncfg(req)

	id, err := s.conncfgSvc.CreateConncfg(ctx, encfg)
	if err != nil {
		return nil, err
	}
	rsp.Id = id

	return &rsp, nil
}
