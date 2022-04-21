package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func komimiSampleCreate() {
	fmt.Println("サンプル追加の処理を行います")
	for i := 0; i < 10; i++ {
		accountInfo := AccountInfo{CustomerId: "0", CustomerPayment: "0", CustomerBalance: "10000"}
		err := accountInfo.Create()
		if err != nil {
			fmt.Printf("i=%v err=%v \n", strconv.Itoa(i), err)
		}
	}
}
func KomimiSampleTruncate() {
	fmt.Println("テーブル初期化の処理を行います")
	accountInfo := AccountInfo{}
	accountInfo.Truncate()
}
func KomimiDeleteAny() {
	fmt.Println("削除したい顧客のidを入力してください>")

	//scan
	scanner.Scan()
	st := scanner.Text()

	//文字列を数字に変換
	i, _ := strconv.Atoi(st)

	//変換した数字を使って検索。エラーは受け取っていない
	accountInfo := AccountInfo{}
	accountInfo.DeleteAny(i)
	return
}
func KomimiSelectAll() {
	fmt.Println("全件抽出します")
	accountInfo := AccountInfo{}
	ac, _ := accountInfo.SelectAll()
	for _, account := range ac {
		fmt.Println(account)
	}
}
func KomimiSelectAny() (accountInfo AccountInfo) {
	fmt.Println("呼び出したい顧客のidを入力してください>")

	//scan
	scanner.Scan()
	id := scanner.Text()

	//変換した数字を使って検索。エラーは受け取っていない
	//ac, _ := accountInfo.SelectAny(id)

	_ = Db.QueryRow("SELECT customer_id FROM account_info WHERE customer_id=$1",
		id,
	).Scan(&accountInfo.CustomerId)

	_ = Db.QueryRow("SELECT customer_balance FROM account_info WHERE customer_id=$1",
		id,
	).Scan(&accountInfo.CustomerBalance)

	fmt.Println(accountInfo)
	return
}
func KomimiUpdateAny() {
	KomimiSelectAll()

	fmt.Println("idを入力してください。")
	ac := KomimiSelectAny()
	fmt.Println()
	ac.CustomerPayment = "10000"
	b, _ := strconv.Atoi(ac.CustomerBalance)
	b = b * 2
	ac.CustomerBalance = strconv.Itoa(b)
	ac.Update()

}

func KomimiImport() {
	fmt.Println("csv読み込み")
	file, err := os.Open("dummy_accounts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	accounts := []AccountInfo{}

	//1行目を飛ばす
	_, err = csvReader.Read()
	if err == io.EOF {
		log.Fatalln(err)
	}

	for {
		line, err := csvReader.Read()
		if err != nil {
			break
		}
		fmt.Println(line)
		customerId := line[0]
		customerBalance := line[1]
		customerPayment := "0"

		if err != nil {
			continue
		}

		accounts = append(accounts, AccountInfo{
			CustomerId:      customerId,
			CustomerPayment: customerPayment,
			CustomerBalance: customerBalance,
		})
	}
	//fmt.Println(accounts)

	for _, a := range accounts {
		a.Import()
	}
}
