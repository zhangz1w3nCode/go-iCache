package user

import (
	"context"
	"fmt"
	"net/http"

	userlogic "visual-state-machine/internal/logic/user"
)

type Service struct {
	user *userlogic.User
}

func NewService(user *userlogic.User) *Service {
	return &Service{
		user: user,
	}
}

func (s *Service) Get(w http.ResponseWriter, r *http.Request) {
	// 设置 CORS 响应头
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// 处理 OPTIONS 预检请求
	if r.Method == "OPTIONS" {
		return
	}

	userDB, err := s.user.GetUser(context.Background(), 1)
	if err != nil {
		return
	}

	fmt.Println(userDB)
}
