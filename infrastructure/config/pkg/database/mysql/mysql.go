package mysql

import (
	"GoFrame/infrastructure/config/pkg/util"
	"GoFrame/infrastructure/po/conncfg"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"GoFrame/global"
)

func NewDB() *gorm.DB {
	return initDB()
}

func initDB() *gorm.DB {
	c := global.Gconfig.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", c.Username, c.Password, util.ParseHost(c.DbHost), c.DbPort, c.DbName, c.Charset)

	cfg := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		// gorm的log设置
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Duration(c.SlowThreshold) * time.Second, // 慢 SQL 阈值
				LogLevel:                  logger.LogLevel(c.LogMode),                   // Log level
				Colorful:                  true,                                         // 彩色打印
				IgnoreRecordNotFoundError: true,                                         // 忽略ErrRecordNotFound（记录未找到）错误
			},
		),
	}
	db, err := gorm.Open(mysql.Open(dsn), cfg)
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(c.MaxOpenConns) //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	sqlDB.SetMaxOpenConns(c.MaxOpenConns) //设置数据库连接池最大连接数
	sqlDB.SetConnMaxLifetime(time.Duration(c.ConnMaxLifetime) * time.Second)

	// 自动创建表
	_ = db.AutoMigrate(
		&po.SysConncfg{},
		&po.SysConncfgLog{},
		&po.SysConncfgCall{},
	)

	// 使用插件
	if err = db.Use(&TracePlugin{}); err != nil {
		panic("initDB:plugin:err")
	}

	return db
}
