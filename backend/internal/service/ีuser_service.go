package service

import (
	"myapp/internal/model"
)

func GetUserInfoLogic(user *model.User, mail string) (error){
	//		Get info from DB
	err := model.GetInfo(user, mail)
	if err != nil {
		return err
	}
	return nil
}

func GetUserBalanceLogic(user *model.User, mail string)error{
	//		Get u_id from DB
	err := model.GetUserByEmail(mail, user)
	if err != nil {
		return err
	}

	//		Get Balance from u_id
	err = model.GetBalance(user)
	if err != nil {
		return err
	}

	return nil
}