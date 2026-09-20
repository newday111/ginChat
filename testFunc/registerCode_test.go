package testfunc

import (
	"fmt"
	"go_jichu/internal/utils"
	"testing"
)

func TestRegisetCode(t *testing.T) {
	code, err := utils.GenerateRegisterCode()
	if err != nil {
		fmt.Println("报错了")
		return
	}
	fmt.Println(code)
}
