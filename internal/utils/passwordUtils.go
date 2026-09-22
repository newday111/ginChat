package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// bcrypt.DefaultCost 默认加密开销（通常为 10，兼顾安全与性能）
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
