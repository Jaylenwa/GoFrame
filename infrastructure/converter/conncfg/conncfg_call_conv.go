package converter

import (
	"GoFrame/infrastructure/po/conncfg"
	"time"

	"gorm.io/plugin/soft_delete"

	entity_conncfg "GoFrame/domain/entity/conncfg"
)

// ConncfgCallConv conncfg_log数据转换
type ConncfgCallConv struct {
}

// E2PConncfgCallAdd entity数据转换成数据库po
func E2PConncfgCallAdd(en *entity_conncfg.ConncfgCall) (poCfg *po.SysConncfgCall) {
	var po po.SysConncfgCall
	po.CreatedAt = time.Now()
	po.UpdatedAt = time.Now()
	po.OpSource = en.ConncfgId

	return &po
}

// E2PConncfgCallDel entity数据转换成数据库po
func E2PConncfgCallDel(en *entity_conncfg.ConncfgCall) *po.SysConncfgCall {
	var po po.SysConncfgCall
	po.UpdatedAt = time.Now()
	po.DeletedAt = soft_delete.DeletedAt(time.Now().Unix())

	return &po
}

// P2EConncfgCall 数据库po转换成entity
func P2EConncfgCall(po *po.SysConncfgCall) *entity_conncfg.ConncfgCall {
	var en entity_conncfg.ConncfgCall
	en.CallId = po.CallId
	en.ConncfgId = po.OpSource
	en.CreatedAt = po.CreatedAt
	en.UpdatedAt = po.UpdatedAt

	return &en
}
