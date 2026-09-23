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
