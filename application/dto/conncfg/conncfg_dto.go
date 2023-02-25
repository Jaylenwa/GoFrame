package dto

// 请求对象
type (
	// CreateConncfgReq 创建conncfg 请求对象
	CreateConncfgReq struct {
		LinkName     string `validate:"required,min=1,max=64" json:"link_name"`
		LinkDesc     string `validate:"omitempty,min=1,max=255" json:"link_desc"`
		ServerName   string `validate:"required,min=1,max=64" json:"server_name"`
		Client       string `validate:"required,min=1,max=64" json:"client"`
		User         string `validate:"required,min=1,max=12" json:"user"`
		Pwd          string `validate:"required,min=1,max=40" json:"pwd"`
		Sid          string `validate:"required,min=1,max=64" json:"sid"`
		Sysnr        string `validate:"required,min=1,max=64" json:"sysnr"`
		CharacterSet string `validate:"required,min=1,max=32" json:"character_set"`
		Language     string `validate:"required,min=1,max=32" json:"language"`
		State        int    `validate:"required,min=1,max=2" json:"state"`
		CfgSource    int    `json:"cfg_source"`
	}
)

// 输出对象
type (
	// CreateConncfgRsp 创建conncfg 返回对象
	CreateConncfgRsp struct {
		Id uint64 `json:"id"`
	}
)
