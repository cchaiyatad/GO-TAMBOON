package controller

import (
	"fmt"
	"os"
	"tamboon/model/transaction"
	"tamboon/service/decrypt"
	"tamboon/service/flag"
	"tamboon/service/payment"
	"time"
)

func App() {
	start := time.Now()
	defer func() { fmt.Printf("Executed time: %s\n", time.Since(start)) }()

	ok := flag.PraseFlag()

	if !ok {
		os.Exit(1)
	}
	fmt.Printf("Performing donations on %s\n", flag.GetFilePath())

	decrypt.Init(flag.GetFilePath())
	defer decrypt.CloseFile()

	decrypt.GetDecrypt()
	payment.Init(flag.GetPublickey(), flag.GetSecretkey())

	tran, _ := transaction.CreateTransaction([]byte("Mr. Holfast J Labingi,3381761,5472068035825145,350,5,2023"))
	payment.Charge(tran)
}
