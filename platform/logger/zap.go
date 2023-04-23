package logger

import (
	"go.uber.org/zap"
)

var (
	Logger      *zap.Logger
	SugarLogger *zap.SugaredLogger
)

func InitLogger() {

	Logger, _ = zap.NewProduction()
	defer Logger.Sync()

	SugarLogger = Logger.Sugar()
	defer SugarLogger.Sync()

}
