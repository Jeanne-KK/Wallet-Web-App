package service

import (
	"myapp/internal/model"
	"github.com/go-playground/validator/v10"
	"errors"
	"fmt"
	"myapp/internal/utils"
)

var validate = validator.New()

func ValidateInput(data model.InputRegister)error{
	err := validate.Struct(data)
	if err != nil {
		return err
	}
	return nil
}

func RegisterUser(data model.InputRegister) (string, error){
	//		Validate input
	err := ValidateInput(data)
	if err != nil {
		return "", fmt.Errorf("validaion fail : %w", err)
	}

	//		Check already use this mail
	row, err := model.CheckMailExist(data.Mail)
	if err != nil {
		return "", fmt.Errorf("internal error: %w", err)
	}
	if row {
		return "", errors.New("email already use")
	}

	//		Hast password
	hash, err := utils.HashPassword(data.Password)
	if err != nil {
		return "", errors.New("fail to hash password")
	}

	//		Save to DB
	newUser := model.InputRegister{
		Mail: data.Mail,
		Password: hash,
		Name: data.Name,
		Surname: data.Surname,
		Phone: data.Phone,
	}
	err = model.CreateUser(newUser)
	if err != nil {
		return "", errors.New("fail to create user")
	}


	//	Create JWT token
	token, err := utils.GenerateToken(newUser.Mail)
	if err != nil {
		return "", errors.New("fail to create token")
	}

	return token, nil
}