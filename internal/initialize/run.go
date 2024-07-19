package initialize

import (
	"fmt"

	"github.com/dinhduyphuoc/go-backend-template/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	InitLogger()
	// InitDatabase()
	r := InitRouter()

	port := global.Config.Server.Port
	global.Logger.Info("Server is running.", zap.Int("port", port))
	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		global.Logger.Error("Server run failed.", zap.Error(err))
	}
}
