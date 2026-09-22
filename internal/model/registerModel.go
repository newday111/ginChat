package model

import (
	"fmt"
	mysqldb "go_jichu/internal/db/mysqlDB"
	userhandler "go_jichu/internal/handlers/userHandler"
)

type user struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Password string `gorm:"size:255;not null" json:"-"`
}

func CreateRegisterUser(registS *userhandler.RegisterUserStruct) (string, error) {
	createRegisterUserMessage := ""
	registerUser := user{
		Email:    registS.Email,
		Password: registS.Password,
	}
	registerResult := mysqldb.GlobalMDB.Create(&registerUser)
	if registerResult.Error != nil {
		createRegisterUserMessage = "create register user failed"
		return createRegisterUserMessage, registerResult.Error
	}
	if registerResult.RowsAffected <= 0 {
		return "", fmt.Errorf("create register user failed, RowsAffected: %d", registerResult.RowsAffected)
	}
	return "", nil
}
