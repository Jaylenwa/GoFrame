package repository

import (
	"context"

	entity_conncfg "GoFrame/domain/entity/conncfg"
)

// ConncfgCallRepository conncfg 调阅历史
//go:generate mockgen --source ./i_conncfg_call_repo.go --destination ./mock/mock_i_conncfg_call_repo.go --package mock
type ConncfgCallRepository interface {
	Create(ctx context.Context, cfg *entity_conncfg.ConncfgCall) (id uint64, err error)
}
