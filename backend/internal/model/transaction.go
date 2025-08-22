package model

import (
	"myapp/internal/db"
	"database/sql"
	"log"
	"fmt"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	T_id int `json:"t_id"`
	W_id int `json:"w_id" validate:"required,numeric"`
	Type string `json:"type" validate:"required,numeric"`
	Amount string `json:"amount" validate:"required,numeric"`
	From_w_id int `json:"from_w_id"`
	To_w_id int `json:"to_w_id" validate:"required,numeric`
	T_create_at string `json:"t_create_at"`
}

func CheckOwnerWallet(tx *sql.Tx, mail string, w_id int) (error) {
	rows, err := tx.Query("select u_mail, u.u_id, w_id from user u, wallet w where u_mail = ? and u.u_id = w.u_id and w_id = ?",mail, w_id)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return nil
	}else{
		return sql.ErrNoRows
	}
}

func CheckToWallet(tx *sql.Tx, t_w_id int)(error){
	rows, err := tx.Query("select w_id from wallet where w_id = ?", t_w_id)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return nil
	}else{
		return sql.ErrNoRows
	}
}

func CheckEnoughBalance(tx *sql.Tx, transaction *Transaction) (error) {
	//		change amount string to decimal
	amount, err := decimal.NewFromString(transaction.Amount)
	if err != nil {
		return fmt.Errorf("invalid amount format")
	}
	//		lock transaction
	var balanceTx string
	err = tx.QueryRow("select w_balance from wallet where w_id = ? for update", transaction.W_id).Scan(&balanceTx)
	if err != nil {
		return err
	}
	err = tx.QueryRow("select w_balance from wallet where w_id = ? for update", transaction.To_w_id).Scan(&balanceTx)
	if err != nil {
		return err
	}

	//		check balance
	var balanceStr string
	err = tx.QueryRow("select w_balance from wallet where w_id = ? and w_balance >= ?", transaction.W_id, amount).Scan(&balanceStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("dont have enought balance")
		}
		return err
	}	

	//		insert transaction in DB
	res, err := tx.Exec("insert into transaction (w_id, type, amount, from_w_id, to_w_id, t_status) values (?, ?, ?, ?, ?, ?)", transaction.W_id, transaction.Type, amount, transaction.W_id, transaction.To_w_id, "0")
	if err != nil {
		return err
	}
	
	InsertId, err := res.LastInsertId()
	if err == nil {
		log.Println(InsertId)
	}

	//		minus balance from owner
	_, err = tx.Exec("update wallet set w_balance=w_balance-? where w_id = ?", amount, transaction.W_id)
	if err != nil {
		return err
	}

	//		plus balance to des
	_, err = tx.Exec("update wallet set w_balance=w_balance+? where w_id = ?", amount, transaction.To_w_id)
	if err != nil {
		return err
	}

	//		update status transaction to success
	_, err = tx.Exec("update transaction set t_status = ? where t_id = ?", "1", InsertId)
	if err != nil {	
		return err
	}
	return nil
}

func Transfer(mail string, transaction *Transaction) (error){
	//		make transaction db
	tx, err := db.DB.Begin()
	if err != nil {
		return fmt.Errorf("cannt create transaction")
	}
	log.Print("make transaction")

	//		check owner wallet with mail
	err = CheckOwnerWallet(tx, mail, transaction.W_id) 	
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("not found wallet in this mail")
	}
	log.Print("have owner wallet")

	//		check des wallet not same tranfer wallet
	if transaction.W_id == transaction.To_w_id {
		return fmt.Errorf("cannt tranfer to same wallet")
	}
	

	//		check have des wallet in db
	err = CheckToWallet(tx, transaction.To_w_id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("dont have des wallet")
	}
	log.Print("have des wallet")

	//		check have enough money and transfer
	err = CheckEnoughBalance(tx, transaction)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf(err.Error())
	}
	log.Print("have enough money")

	//		commit transaction db
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("Commit fail")
	}
	log.Println("success transfer")

	return nil
}