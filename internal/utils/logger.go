package middleware

import (
	"fmt"
	"go_jichu/conf"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	AccessLog *zap.Logger
	ErrorLog  *zap.Logger
)

func InitLogger(cfg *conf.ConfigStruct) error {
	// 创建 logs 目录
	fmt.Println(cfg.App.LogPath)
	if err := os.MkdirAll(cfg.App.LogPath, 0755); err != nil {
		return err
	}

	accessPath := cfg.App.LogPath + "/access-%Y-%m-%d.log"
	errorPath := cfg.App.LogPath + "/error-%Y-%m-%d.log"

	// access 日志
	accessWriter, err := rotatelogs.New(
		accessPath,
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithMaxAge(30*24*time.Hour),
	)
	if err != nil {
		return err
	}

	// error 日志
	errorWriter, err := rotatelogs.New(
		errorPath,
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithMaxAge(30*24*time.Hour),
	)
	if err != nil {
		return err
	}

	// JSON 格式
	encoderConfig := zap.NewProductionEncoderConfig()

	// 时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	// Access Logger
	accessCore := zapcore.NewCore(
		encoder,
		zapcore.AddSync(accessWriter),
		zap.InfoLevel,
	)

	AccessLog = zap.New(accessCore)

	// Error Logger
	errorCore := zapcore.NewCore(
		encoder,
		zapcore.AddSync(errorWriter),
		zap.ErrorLevel,
	)

	ErrorLog = zap.New(errorCore)

	return nil
}

// Sync 刷新日志
func Sync() {
	if AccessLog != nil {
		_ = AccessLog.Sync()
	}

	if ErrorLog != nil {
		_ = ErrorLog.Sync()
	}
}
