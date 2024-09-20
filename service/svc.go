package service

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormless "gorm.io/gorm/logger"
	syslog "log"
	"os"
	"time"
)

var dbConn *gorm.DB

func InitCon() {
	dbConn = NewGorm("root:root_password@tcp(localhost:3306)/my_database", "off")
}

func NewGorm(sourceUrl, logSwitch string) *gorm.DB {
	// gorm config
	gcfg := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		AllowGlobalUpdate:                        false,
	}

	// 开启日志
	if logSwitch == "on" {
		gcfg.Logger = gormless.New(
			syslog.New(os.Stdout, "\r\n", syslog.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容）
			gormless.Config{
				SlowThreshold:             time.Second,   // 慢 SQL 阈值
				LogLevel:                  gormless.Info, // 日志级别 Debug()
				IgnoreRecordNotFoundError: false,         // 不忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  true,          // 使用彩色打印
			},
		)
	}

	db, err := gorm.Open(mysql.Open(sourceUrl), gcfg)
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(40)
	sqlDB.SetMaxOpenConns(40)

	return db
}
