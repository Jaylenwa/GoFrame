package service

import (
	"context"

	entity_conncfg "GoFrame/domain/entity/conncfg"
)

// ConncfgService conncfg service 接口
//go:generate mockgen --source ./conncfg_agg.go --destination ./mock/mock_conncfg_agg.go --package mock
type ConncfgService interface {
	CreateConncfg(ctx context.Context, cfgEntity *entity_conncfg.Conncfg) (id uint64, err error)
}
