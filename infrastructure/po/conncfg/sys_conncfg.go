package po

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

// SysConncfg conncfg配置
type SysConncfg struct {
	ConncfgId    uint64                `gorm:"primary_key;auto_increment" json:"conncfg_id"`
	CreatedAt    time.Time             `gorm:"column:created_at;type:datetime;" json:"created_at"` // 创建时间
	UpdatedAt    time.Time             `gorm:"column:updated_at;type:datetime;" json:"updated_at"` // 修改时间
	DeletedAt    soft_delete.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`               // 删除时间
	CreatedBy    string                `gorm:"type:varchar(64); not null;" json:"created_by"`
	UpdatedBy    string                `gorm:"size:64;not null;" json:"updated_by"`
	DeletedBy    string                `gorm:"size:64;not null;" json:"deleted_by"`
	LinkName     string                `gorm:"size:64;not null;" json:"link_name"`             // 链接名
	LinkDesc     string                `gorm:"size:255;not null;" json:"link_desc"`            // 链接描述
	ServerName   string                `gorm:"size:64;not null;" json:"server_name"`           // sap 服务名
	Client       string                `gorm:"size:32;not null;" json:"client"`                // sap client
	User         string                `gorm:"size:32;not null;" json:"user"`                  // sap user
	Pwd          string                `gorm:"size:128;not null;" json:"pwd"`                  // sap pwd
	Sid          string                `gorm:"size:64;not null;" json:"sid"`                   // sap sid
	Sysnr        string                `gorm:"size:64;not null;" json:"sysnr"`                 // sap sysnr
	CharacterSet string                `gorm:"size:32;not null;" json:"character_set"`         // sap 字符集
	Language     string                `gorm:"size:32;not null;" json:"language"`              // sap 语言
	CfgSource    int                   `gorm:"not null;" json:"cfg_source"`                    // 配置源：1:sap
	State        int                   `gorm:"type:int(10);not null;default: 2;" json:"state"` // 1启用2不启用
}

// SysConncfgLog 配置日志
type SysConncfgLog struct {
	LogId     uint64                `gorm:"primary_key;auto_increment" json:"log_id"`
	CreatedAt time.Time             `gorm:"column:created_at;type:datetime;" json:"created_at"` // 创建时间
	UpdatedAt time.Time             `gorm:"column:updated_at;type:datetime;" json:"updated_at"` // 修改时间
	DeletedAt soft_delete.DeletedAt `gorm:"default:NULL" json:"deleted_at"`
	CreatedBy string                `gorm:"size:64;not null;" json:"created_by"`
	UpdatedBy string                `gorm:"size:64;not null;" json:"updated_by"`
	DeletedBy string                `gorm:"size:64;not null;" json:"deleted_by"`
	OpType    string                `gorm:"size:64;not null;" json:"op_type"`   // 操作类型：create:创建;update:修改;delete:删除
	OpSource  uint64                `gorm:"size:64;not null;" json:"op_source"` // 操作资源，conncfg_id
	OpDetail  string                `gorm:"size:64;not null;" json:"op_detail"` // 操作具体内容
}

// SysConncfgCall SysConncfg 调阅历史
type SysConncfgCall struct {
	CallId    uint64                `gorm:"primary_key;auto_increment" json:"call_id"`
	CreatedAt time.Time             `gorm:"column:created_at;type:datetime;" json:"created_at"` // 创建时间
	UpdatedAt time.Time             `gorm:"column:updated_at;type:datetime;" json:"updated_at"` // 修改时间
	DeletedAt soft_delete.DeletedAt `gorm:"default:NULL" json:"deleted_at"`
	CreatedBy string                `gorm:"type:varchar(64); not null;" json:"created_by"`
	UpdatedBy string                `gorm:"size:64;not null;" json:"updated_by"`
	DeletedBy string                `gorm:"size:64;not null;" json:"deleted_by"`
	OpSource  uint64                `gorm:"size:64;not null;" json:"op_source"`
}
