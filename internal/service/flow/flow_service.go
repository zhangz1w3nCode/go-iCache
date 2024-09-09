package flow

import (
	"context"
)

type Service struct {
}

func NewFlowService() *Service {
	return &Service{}
}

func (s *Service) CreateFlow(ctx context.Context) error {

	return nil
}
