package entity

import "time"

// Conncfg  entity
type Conncfg struct {
	ConncfgId    uint64     `json:"conncfg_id"`
	LinkName     string     `json:"link_name"`     // 连接名
	LinkDesc     string     `json:"link_desc"`     // 连接描述
	ServerName   string     `json:"server_name"`   // 服务名
	Client       string     `json:"client"`        // sap client
	User         string     `json:"user"`          // user
	Pwd          string     `json:"pwd"`           // pwd
	Sid          string     `json:"sid"`           // sap sid
	Sysnr        string     `json:"sysnr"`         // sap sysnr
	CharacterSet string     `json:"character_set"` // 字符集
	Language     string     `json:"language"`      // 语言
	CfgSource    int        `json:"cfg_source"`    // 配置源(1:sap)
	CreatedAt    time.Time  `json:"created_at"`    // 创建时间
	UpdatedAt    time.Time  `json:"updated_at"`    // 修改时间
	DeletedAt    *time.Time `json:"deleted_at"`    // 删除时间
	CreatedBy    string     `json:"created_by"`    // 创建者
	UpdatedBy    string     `json:"updated_by"`    // 修改者
	DeletedBy    string     `json:"deleted_by"`    // 删除者
	State        int        `json:"state"`         // 1启用2不启用
}

// ConncfgLog entity cfg操作日志
type ConncfgLog struct {
	LogId     uint64     `json:"log_id"`
	OpType    string     `json:"op_type"`    // 操作类型（create:创建、update:更新、delete:删除）
	OpSource  uint64     `json:"op_source"`  // 操作资源ID
	OpDetail  string     `json:"op_detail"`  // 操作详情
	CreatedAt time.Time  `json:"created_at"` // 创建时间
	UpdatedAt time.Time  `json:"updated_at"` // 修改时间
	DeletedAt *time.Time `json:"deleted_at"` // 删除时间
	CreatedBy string     `json:"created_by"` // 创建者
	UpdatedBy string     `json:"updated_by"` // 修改者
	DeletedBy string     `json:"deleted_by"` // 删除者
}

// ConncfgCall entity 调用历史
type ConncfgCall struct {
	CallId    uint64     `json:"call_id"`
	ConncfgId uint64     `json:"conncfg_id"`
	CreatedAt time.Time  `json:"created_at"` // 创建时间
	UpdatedAt time.Time  `json:"updated_at"` // 修改时间
	DeletedAt *time.Time `json:"deleted_at"` // 删除时间
}

// ConncfgCheck entity cfg校验
type ConncfgCheck struct {
	LinkName     string `json:"link_name"`     // 连接名
	LinkDesc     string `json:"link_desc"`     // 连接描述
	ServerName   string `json:"server_name"`   // 服务名
	Client       string `json:"client"`        // sap client
	User         string `json:"user"`          // user
	Pwd          string `json:"pwd"`           // pwd
	Sid          string `json:"sid"`           // sap sid
	Sysnr        string `json:"sysnr"`         // sap sysnr
	CharacterSet string `json:"character_set"` // 字符集
	Language     string `json:"language"`      // 语言
}
