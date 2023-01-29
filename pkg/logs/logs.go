package logs

import (
	"go.uber.org/zap"
)

type Logs struct {
	Logger *zap.Logger
}

func New() (*Logs, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return &Logs{
		Logger: logger,
	}, nil
}
