package config

import (
	"fmt"
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v3"
)

// ServerConf 服务端口，名称，mode
type ServerConf struct {
	Lang        string `json:"lang"`
	Host        string `json:"host"`
	PublicPort  string `yaml:"public_port"`
	PrivatePort string `yaml:"private_port"`
	ServerName  string `yaml:"server_name"`
	Mode        string `yaml:"mode"`
	IsDebug     string `yaml:"is_debug"` // 是否开启从本地读取SAP数据
}

// MysqlConf mysql
type MysqlConf struct {
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	DbHost          string `yaml:"db_host"`
	DbPort          int    `yaml:"db_port"`
	DbName          string `yaml:"db_name"`
	Charset         string `yaml:"charset"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
	LogMode         int    `yaml:"log_mode"`
	SlowThreshold   int    `yaml:"slow_threshold"`
}

// LogConf 日志
type LogConf struct {
	LogFileDir string `yaml:"log_file_dir"`
	AppName    string `yaml:"app_name"`
	MaxSize    int    `yaml:"max_size"`    //文件多大开始切分
	MaxBackups int    `yaml:"max_backups"` //保留文件个数
	MaxAge     int    `yaml:"max_age"`     //文件保留最大实际
	LogLevel   string `json:"log_Level"`   // 日志级别
}

// NsqConf nsq
type NsqConf struct {
	NsqProducerHost  string `yaml:"nsq_producer_host"`
	NsqProducerPort  string `yaml:"nsq_producer_port"`
	NsqSubscribeHost string `yaml:"nsq_subscribe_host"`
	NsqSubscribePort string `yaml:"nsq_subscribe_port"`
}

type Config struct {
	Server ServerConf
	Mysql  MysqlConf
	Log    LogConf
	Nsq    NsqConf
}

func NewConfig() (conf *Config) {
	conf = initConfig()
	return
}

var (
	configOnce sync.Once
	config     *Config
)

// InitConfig 读取配置
func initConfig() *Config {
	configOnce.Do(func() {
		fileBytes, err := ioutil.ReadFile("infrastructure/config/gf.yaml")
		if err != nil {
			panic(fmt.Sprintf("load gf.yaml failed: %v", err))
		}

		err = yaml.Unmarshal(fileBytes, &config)
		if err != nil {
			panic(fmt.Sprintf("unmarshal yaml file failed: %v", err))
		}
	})

	return config
}
