package services

import (
	"fmt"

	"../config"
	"../dao"
	"../models"
	"github.com/joaopandolfi/blackwhale/utils"
)

type UserService interface {
	Login(username string, password string) (user models.User, success bool, err error)
	NewUserClient(user models.User) (result models.User, err error)
	NewUser(user models.User) (result models.User, err error)
	CheckToken(userid int, token string) (success bool, err error)
}

type User struct {
	UserDAO dao.UserDAO
}

func (cc User) CheckToken(userid int, token string) (success bool, err error) {
	_, success, err = cc.UserDAO.CheckToken(models.User{Token: token, ID: userid})
	return
}

func (cc User) Login(username string, password string) (user models.User, success bool, err error) {
	return cc.UserDAO.Login(models.User{Username: username, Password: password})
}

// New basic client user
func (cc User) NewUserClient(user models.User) (result models.User, err error) {
	user.Level = models.USER_CLIENT
	return cc.NewUser(user)
}

// NewUser Generic
func (cc User) NewUser(user models.User) (result models.User, err error) {
	user.Password, err = utils.HashPassword(user.Password)
	user.Token, err = utils.HashPassword(fmt.Sprintf(config.Config.Token, user.Username))

	if err != nil {
		user.Password = ""
		return
	}
	return cc.UserDAO.NewUser(user)
}
