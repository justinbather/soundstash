package main

import (
	"fmt"

	"go.uber.org/zap"
)

func newLogger() *zap.SugaredLogger {
	log, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("Error creating logger %v", err)
	}

	return log.Sugar()
}
