package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/dinhduyphuoc/go-backend-template/global"
	"github.com/dinhduyphuoc/go-backend-template/pkg/settings"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	// DPanic, Panic and Fatal level can not be set by user
	DebugLevelStr   string = "debug"
	InfoLevelStr    string = "info"
	WarningLevelStr string = "warning"
	ErrorLevelStr   string = "error"
)

func FormatLogFileName(fileName string, path string) string {
	datetime := time.Now().Format("2006-01-02")
	return fmt.Sprintf("%s/%s/%s-%s.log", path, datetime, fileName, datetime)
}

func InitLogger(settings settings.LoggerConfig) error {
	var level zapcore.Level
	switch settings.LogLevel {
	case DebugLevelStr:
		level = zap.DebugLevel
	case InfoLevelStr:
		level = zap.InfoLevel
	case WarningLevelStr:
		level = zap.WarnLevel
	case ErrorLevelStr:
		level = zap.ErrorLevel
	default:
		return fmt.Errorf("unknown log level '%s'", settings.LogLevel)
	}

	fileName := FormatLogFileName(settings.FileName, settings.Path)

	ws := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    settings.MaxSize, // MB
		MaxBackups: settings.MaxBackups,
		MaxAge:     settings.MaxAge, // days
		Compress:   settings.Compress,
	})
	core := zapcore.NewCore(
		GetLogEncoder(),
		// write to stdout as well as log files
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), ws),
		zap.NewAtomicLevelAt(level),
	)
	if settings.DevMode {
		global.Logger = zap.New(core, zap.AddCaller(), zap.Development())
	} else {
		global.Logger = zap.New(core)
	}
	zap.ReplaceGlobals(global.Logger)
	return nil
}

func GetLogEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	// Format timestamp to ISO8601
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Format keyword ts to Time
	encoderConfig.TimeKey = "ts"

	// Format keyword level to uppercase
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}
