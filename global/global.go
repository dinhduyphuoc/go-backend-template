package global

import (
	"github.com/dinhduyphuoc/go-backend-template/pkg/settings"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config   settings.Config
	Logger   *zap.Logger
	Database *gorm.DB
)
