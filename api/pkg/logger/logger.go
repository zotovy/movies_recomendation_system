package logger

import (
	"app/pkg/config"
	"fmt"
	"go.uber.org/zap"
	"os"
	"time"
)

type Logger = zap.Logger

func GetLogger(config *config.Config) (*zap.Logger, error) {
	if config.Environment == "release" {
		filename := fmt.Sprintf("../logs/%s.log", time.Now().Format(time.RFC822))

		_, err := os.Create(filename)
		if err != nil {
			println(err)
			return nil, err
		}

		cfg := zap.NewProductionConfig()
		cfg.OutputPaths = []string{filename}
		cfg.ErrorOutputPaths = []string{filename}

		logger, _ := cfg.Build()
		return logger, nil
	}

	return zap.NewDevelopment()
}
