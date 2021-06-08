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

func beginTransaction() {
	producer, filePointer := decrypt.GetProducer(flag.GetFilePath())

	for line := range producer {
		// fmt.Println(line)
		if line == nil {
			// EOF
			break
		}

		tran, err := transaction.CreateTransaction(line)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(tran)
	}

	decrypt.CleanProducer(filePointer)
}

func App() {
	start := time.Now()
	defer func() { fmt.Printf("Executed time: %s\n", time.Since(start)) }()

	if ok := flag.PraseFlag(); !ok {
		// TODO: print error
		os.Exit(1)
	}

	payment.Init(flag.GetPublickey(), flag.GetSecretkey(), flag.GetNumberTask())

	fmt.Printf("Performing donations on %s\n", flag.GetFilePath())

	beginTransaction()

	// payment.Run(producer)
}
