package controllers

import (
	"github.com/joaopandolfi/hydra-back/dao"
	"github.com/joaopandolfi/hydra-back/services"
)

// NewUserService - Factory
func NewUserService() services.UserService {
	return services.User{
		UserDAO: dao.User{},
	}
}
