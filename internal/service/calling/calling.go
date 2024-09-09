package calling

import (
	"context"
	"fmt"
	"visual-state-machine/internal/logic/calling"
)

type Service struct {
	user *calling.UserLogic
}

func NewCallingService(user *calling.UserLogic) *Service {
	return &Service{
		user: user,
	}
}

func (s *Service) GetUser(ctx context.Context) error {
	userDB, err := s.user.GetUser(ctx, 1)
	if err != nil {
		return err
	}
	fmt.Println(userDB)
	return nil
}
