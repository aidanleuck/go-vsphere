package common

import (
	"log"

	"go.uber.org/zap"
)

func initLogger() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Can't initialize zap logger: %v", err)
	}
	sugar := logger.Sugar()
	defer sugar.Sync()
	return sugar
}
