package service

import (
	"myapp/internal/model"
	"errors"
	"log"
	"github.com/shopspring/decimal"
)

func BeforeTransferLogic(transaction *model.Transaction, mail string)(interface{}, string, error){

	//		Check type of transaction
	if transaction.Type == "1" {

		//		Get UserOwn and UserDes
		var data model.CheckUser	
		err := model.CheckDataTransfer(mail, transaction.W_id, transaction.Type, transaction.To_w_id, &data)
		if err != nil {
			log.Println(err)
			return nil, "", err
		}

		return data, transaction.Type ,nil	
	}else{
		return nil, "", errors.New("Dont have this action")
	}
}

func TransferLogic(transaction *model.Transaction, mail string) (string, error) {
	//		validate input
	err := validate.Struct(transaction)
	if err != nil{
		log.Println(err)
		return "", err
	}

	//		String to number		
	amount, err := decimal.NewFromString(transaction.Amount)
	if err != nil {
		log.Println(err)
		return "", err
	}
	
	//		Amount cannt < 1
	if amount.Cmp(decimal.NewFromInt(1)) < 0 {
		return "", errors.New("Amount < 0")
	}

	//		Check type transaction
	if transaction.Type == "1" {
		err := model.Transfer(mail, transaction)
		if err != nil {
			log.Println(err)
			return "", err
		}

		return "1", nil
	}else{
		return "", errors.New("Dont have action")
	}

}