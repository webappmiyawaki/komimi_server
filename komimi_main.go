package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type AccountInfo struct {
	CustomerId                 string
	CustomerPayment            string
	CustomerBalance            string
	CustomerPaymentInformation bool
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=postgres dbname=komimi_server password=test sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func main() {
	var in string

	fmt.Println("select (1:create 2:truncate 3: selectAll 4: selectAny 5: update 8: importCSV　9:server)")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		in = scanner.Text()
		fmt.Println("in: ", in)
		switch in {
		case "1":
			komimiSampleCreate()
			goto L
		case "2":
			KomimiSampleTruncate()
			goto L
		case "3":
			KomimiSelectAll()
			goto L
		case "4":
			KomimiSelectAny()
			goto L
		case "5":
			KomimiUpdateAny()
			goto L
		case "6":
			KomimiDeleteAny()
			goto L
		case "8":
			KomimiImport()
			goto L
		case "9":
			komimiServer()
			goto L
		default:
			fmt.Println("もう一度入力してください。")
			continue
		}
	}
L:
}
