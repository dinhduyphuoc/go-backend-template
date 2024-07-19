package initialize

import (
	"fmt"

	"github.com/dinhduyphuoc/go-backend-template/global"
	"github.com/dinhduyphuoc/go-backend-template/pkg/logger"
	"go.uber.org/zap"
)

func InitLogger() {
	settings := global.Config.Logger
	err := logger.InitLogger(settings)
	if err != nil {
		panic(fmt.Errorf("fatal error initializing logger: %s", err))
	}

	global.Logger.Info("Logger initialized successfully", zap.String("level", settings.LogLevel))
}
