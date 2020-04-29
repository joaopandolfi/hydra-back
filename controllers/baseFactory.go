package controllers

import (
"../dao"
"../services"
)

// NewUserService - Factory
func NewUserService() services.UserService {
	return services.User{
		UserDAO: dao.User{},
	}
}
