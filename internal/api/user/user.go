package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	userlogic "visual-state-machine/internal/logic/user"
)

type Api struct {
	user *userlogic.User
}

func New(user *userlogic.User) *Api {
	return &Api{
		user: user,
	}
}

func (s *Api) Get(c *gin.Context) {
	id := c.Param("id")
	userDB, err := s.user.GetUser(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "user not found",
		})
	}
	c.JSON(http.StatusOK, userDB)
}

func (s *Api) List(c *gin.Context) {
	users, err := s.user.List(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "users not found",
		})
	}
	c.JSON(http.StatusOK, users)
}
