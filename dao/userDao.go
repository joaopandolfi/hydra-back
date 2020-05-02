package dao

import (
	"fmt"

	"../models"
	"github.com/joaopandolfi/blackwhale/remotes/mongo"
	"github.com/joaopandolfi/blackwhale/utils"
	"golang.org/x/xerrors"
	"gopkg.in/mgo.v2/bson"
)

// Dao responsável por gerir os usários

type UserDAO interface {
	NewUser(user models.User) (result models.User, err error)
	Login(user models.User) (result models.User, success bool, err error)
	CheckToken(user models.User) (result models.User, success bool, err error)
}

type User struct {
}

// NewUser -
// Create new user on database
func (cc User) NewUser(user models.User) (result models.User, err error) {
	//Arrange
	session, err := mongo.NewSession()
	if err != nil {
		err = xerrors.Errorf("Unable to connect to mongo: %v", err)
		//utils.CriticalError("Unable to connect to mongo:", err.Error())
		return
	}
	id := mongo.GetNextID("user_id")
	user.ID = id
	user.Token, err = utils.NewJwtToken(utils.Token{
		ID:          fmt.Sprint(user.ID),
		Institution: fmt.Sprint(user.Instution),
		Permission:  fmt.Sprint(user.Level),
	}, 60)

	result = user

	err = session.GetCollection("user").Insert(&user)
	if err != nil {
		err = xerrors.Errorf("Insert user error %v", err)
	}
	return
}

// CheckToken logged user
func (cc User) CheckToken(user models.User) (result models.User, success bool, err error) {
	var results []models.User
	session, err := mongo.NewSession()

	if err != nil {
		err = xerrors.Errorf("Unable to connect to mongo: %v", err)
		return
	}

	err = session.GetCollection("user").Find(bson.M{"id": user.ID, "token": user.Token}).All(&results)
	if err != nil {
		err = xerrors.Errorf("query user error %v", err)
		return
	}

	//utils.Debug("Check Token result [id,token,result]",user.ID,user.Token,results)
	if len(results) > 0 {
		success = true
		result = results[0]
	}

	return
}

// Login based on hashed user
func (cc User) Login(user models.User) (result models.User, success bool, err error) {
	var users []models.User
	session, err := mongo.NewSession()
	if err != nil {
		err = xerrors.Errorf("Unable to connect to mongo: %v", err)
		return
	}

	err = session.GetCollection("user").Find(bson.M{"username": user.Username, "instution": user.Instution}).All(&users)
	if err != nil {
		err = xerrors.Errorf("query user error %v", err)
		return
	}

	if len(users) > 0 {
		if utils.CheckPasswordHash(user.Password, users[0].Password) {
			result = users[0]
			result.Password = ""
			success = true
			return
		}
	}
	return
}
