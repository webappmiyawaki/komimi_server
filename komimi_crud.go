package main

import (
	"fmt"
)

func (accountInfo *AccountInfo) Create() (err error) {
	fmt.Printf("id:%v name:%v balance:%v\n", accountInfo.CustomerId, accountInfo.CustomerPayment, accountInfo.CustomerBalance)

	Db.Exec("INSERT INTO account_info(customer_id, customer_name, customer_balance) VALUES ($1,$2,$3)",
		accountInfo.CustomerId,
		accountInfo.CustomerPayment,
		accountInfo.CustomerBalance,
	)
	return
}
func (accountInfo *AccountInfo) Import() (err error) {
	Db.Exec("INSERT INTO account_info(customer_id, customer_balance, customer_payment ) VALUES ($1,$2,$3)",
		accountInfo.CustomerId,
		accountInfo.CustomerBalance,
		accountInfo.CustomerPayment,
	)
	return
}

func (accountInfo *AccountInfo) Truncate() (err error) {
	Db.Exec("TRUNCATE TABLE account_info")
	return
}

func (accountInfo *AccountInfo) SelectAll() (ac []AccountInfo, err error) {
	rows, _ := Db.Query("SELECT * FROM account_info ORDER BY customer_id asc")

	for rows.Next() {
		var a AccountInfo
		rows.Scan(&a.CustomerId, &a.CustomerPayment, &a.CustomerBalance)
		ac = append(ac, a)
	}
	rows.Close()
	return
}

func (accountInfo *AccountInfo) DeleteAny(id int) (err error) {
	Db.Exec("DELETE FROM account_info WHERE customer_id=$1",
		id,
	)
	return
}

func (accountInfo *AccountInfo) SelectAny(id int) (ac []AccountInfo, err error) {
	rows, _ := Db.Query("SELECT * FROM account_info WHERE customer_id=$1",
		id,
	)

	for rows.Next() {
		var a AccountInfo
		rows.Scan(&a.CustomerId, &a.CustomerPayment, &a.CustomerBalance)
		ac = append(ac, a)
	}
	rows.Close()
	return
}

func (accountInfo *AccountInfo) SelectAnyId() (ac []AccountInfo) {
	rows, _ := Db.Query("SELECT * FROM account_info WHERE customer_id=$1",
		accountInfo.CustomerId,
	)

	for rows.Next() {
		var a AccountInfo
		rows.Scan(&a.CustomerId, &a.CustomerPayment, &a.CustomerBalance)
		ac = append(ac, a)
	}
	rows.Close()
	return
}

func (accountInfo *AccountInfo) Update() (ac []AccountInfo, err error) {
	Db.Exec("UPDATE account_info SET customer_name=$1,customer_balance=$2 WHERE customer_id=$3",
		accountInfo.CustomerPayment,
		accountInfo.CustomerBalance,
		accountInfo.CustomerId,
	)
	return
}
