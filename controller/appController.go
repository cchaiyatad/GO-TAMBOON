package controller

import (
	"fmt"
	"os"
	"tamboon/service/decrypt"
	"tamboon/service/flag"
	"tamboon/service/payment"
	"time"
)

func get() {
	prod, fp := decrypt.GetProducer(flag.GetFilePath())

	for line := range prod {
		fmt.Println(line)
	}

	decrypt.CleanProducer(fp)
}

func App() {
	start := time.Now()
	defer func() { fmt.Printf("Executed time: %s\n", time.Since(start)) }()

	ok := flag.PraseFlag()

	if !ok {
		// TODO: print error
		os.Exit(1)
	}

	payment.Init(flag.GetPublickey(), flag.GetSecretkey(), flag.GetNumberTask())

	fmt.Printf("Performing donations on %s\n", flag.GetFilePath())
	// prod := decrypt.GetProducer(flag.GetFilePath())

	// for line := range prod {
	// 	fmt.Println(line)
	// }
	go get()
	time.Sleep(time.Second * 20)
	// decrypt.Init(flag.GetFilePath())
	// defer decrypt.CloseFile()
	// go get()
	// prod := decrypt.Producer()

	// for line := range prod {
	// 	fmt.Println(line)
	// }
	// payment.Run(prod)
}
