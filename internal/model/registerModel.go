package model

import (
	"fmt"
	mysqldb "go_jichu/internal/db/mysqlDB"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Password string `gorm:"size:255;not null" json:"-"`
}

type LoginUser struct {
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

func CreateRegisterUser(user *User) (string, error) {
	createRegisterUserMessage := ""
	registerResult := mysqldb.GlobalMDB.Create(user)
	if registerResult.Error != nil {
		createRegisterUserMessage = "create register user failed"
		return createRegisterUserMessage, registerResult.Error
	}
	if registerResult.RowsAffected <= 0 {
		return "", fmt.Errorf("create register user failed, RowsAffected: %d", registerResult.RowsAffected)
	}
	return "", nil
}

func QueryUserInfoToEmail(loginU *LoginUser) (*LoginUser, error) {
	var loginUser LoginUser
	err := mysqldb.GlobalMDB.Table("users").Where("email = ?", loginU.Email).First(&loginUser).Error
	if err != nil {
		return nil, err
	}
	return &loginUser, nil
}
