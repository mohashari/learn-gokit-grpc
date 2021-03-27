package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	logger log.Logger
}

type Service interface {
	Add(ctx context.Context, numA, numB float32) (float32, error)
}

func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (s *service) Add(ctx context.Context, numA, numB float32) (float32, error) {
	level.Info(s.logger).Log("method", "Add", "numA", numA, "numB", numB)
	return numA + numB, nil
}
