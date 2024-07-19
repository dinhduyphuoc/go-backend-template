package initialize

import (
	"fmt"
	"time"

	"github.com/dinhduyphuoc/trieuhao-server-admin/global"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CheckConnectionError(err error, errorString string) {
	if err != nil {
		global.Logger.Error(errorString, zap.Error(err))
	}
}

func InitDatabase() {
	settings := global.Config.Database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		settings.Host,
		settings.User,
		settings.Password,
		settings.DbName,
		settings.Port,
		settings.SslMode,
		settings.TimeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	CheckConnectionError(err, "Failed connecting to database.")

	global.Logger.Info("Database connected successfully.")
	global.Database = db

	// Setup database connection pool
	SetupPool()
}

func SetupPool() {
	settings := global.Config.Database
	db, err := global.Database.DB()
	if err != nil {
		global.Logger.Error("Failed to get database connection pool.", zap.Error(err))
	}
	db.SetConnMaxIdleTime(time.Duration(settings.MaxIdleTime) * time.Second)
	db.SetConnMaxLifetime(time.Duration(settings.ConMaxLifetime) * time.Second)
	db.SetMaxIdleConns(settings.MaxIdleConns)
	db.SetMaxOpenConns(settings.MaxOpenConns)
}
