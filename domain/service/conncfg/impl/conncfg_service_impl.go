package impl

import (
	service "GoFrame/domain/service/conncfg"
	repository "GoFrame/infrastructure/repository/conncfg"
	"GoFrame/infrastructure/repository/conncfg/impl"
	"context"
	"sync"

	entity_conncfg "GoFrame/domain/entity/conncfg"
)

var (
	conncfgServiceImplOnce sync.Once
	conncfgServiceImpl     service.ConncfgService
)

type conncfgService struct {
	cfgCallRepo repository.ConncfgCallRepository
}

func NewConncfgServiceImpl() service.ConncfgService {
	conncfgServiceImplOnce.Do(func() {
		conncfgServiceImpl = &conncfgService{
			cfgCallRepo: impl.NewconncfgCallRepoImpl(),
		}
	})
	return conncfgServiceImpl
}

// CreateConncfg 增加配置
func (a *conncfgService) CreateConncfg(ctx context.Context, cfgEntity *entity_conncfg.Conncfg) (id uint64, err error) {

	var cfg entity_conncfg.ConncfgCall
	cfg.CallId = 1
	_, _ = a.cfgCallRepo.Create(ctx, &cfg)

	return
}
