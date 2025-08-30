package logger

import (
	"go.uber.org/zap"
)

func NewLogger(production bool) *zap.Logger {
	var log *zap.Logger
	if production {
		log, _ = zap.NewProduction()
	} else {
		log, _ = zap.NewDevelopment()
	}

	return log
}
