package service

import (
	"myapp/internal/model"
	"errors"
	"fmt"
	"myapp/internal/utils"
)

func ValidateInputLogin(data model.InputLogin) error {
	err := validate.Struct(data)
	if err != nil {
		return err
	}
	return nil

}

func LoginUser(data model.InputLogin) (string, error) {
	//		Validate Input
	err := ValidateInputLogin(data)
	if err != nil {
		return "", fmt.Errorf("validaion fail : %w", err)
	}

	//		Check user from mail
	row, err := model.GetPassHash(data.Mail)
	if err != nil {
		return "", fmt.Errorf("Invalid email")
	}

	//		Check Pass
	if !utils.CheckPassword(data.Password, row) {
		return "", fmt.Errorf("Invalid password")
	}

	//		Create JWT token
	token, err := utils.GenerateToken(data.Mail)
	if err != nil {
		return "", errors.New("fail to create token")
	}

	return token, nil


}