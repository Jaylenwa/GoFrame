package impl

import (
	"GoFrame/infrastructure/converter/conncfg"
	conncfgCallRepository "GoFrame/infrastructure/repository/conncfg"
	"context"
	"sync"

	"gorm.io/gorm"

	entity_conncfg "GoFrame/domain/entity/conncfg"
	"GoFrame/global"
)

var (
	conncfgCallRepoImplOnce sync.Once
	conncfgCallRepoImpl     conncfgCallRepository.ConncfgCallRepository
)

type conncfgCallRepo struct {
	db *gorm.DB
}

func NewconncfgCallRepoImpl() conncfgCallRepository.ConncfgCallRepository {
	conncfgCallRepoImplOnce.Do(func() {
		conncfgCallRepoImpl = &conncfgCallRepo{
			db: global.GDB,
		}
	})
	return conncfgCallRepoImpl
}

func (r *conncfgCallRepo) Create(ctx context.Context, cfg *entity_conncfg.ConncfgCall) (id uint64, err error) {
	po := converter.E2PConncfgCallAdd(cfg)
	err = r.db.Create(&po).Error

	return
}
