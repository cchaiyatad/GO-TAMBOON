package controller

import (
	"fmt"
	"os"
	"tamboon/model/summary"
	"tamboon/model/transaction"
	"tamboon/service/decrypt"
	"tamboon/service/flag"
	"tamboon/service/payment"
	"tamboon/service/summaries"
	"time"

	"github.com/omise/omise-go"
)

func beginTransaction(client *omise.Client, consumers chan *summary.Summary) {
	producer, filePointer, err := decrypt.GetProducer(flag.GetFilePath())
	if err != nil {
		//TODO: Error
		decrypt.CleanProducer(filePointer)
		// fmt.Println(err)
		os.Exit(1)
	}

	for line := range producer {
		// fmt.Println(line)
		if line == nil {
			// EOF
			break
		}

		tran, err := transaction.CreateTransaction(line)
		if err != nil {
			// fmt.Println(err)
			continue
		}
		// fmt.Println(tran)
		// err = payment.BeginCharge(tran, client)
		consumer := <-consumers
		consumer.Update(*tran, err == nil)
		consumers <- consumer
	}
	fmt.Println("Done.")
	decrypt.CleanProducer(filePointer)
	summaries.CleanConsumer(consumers)
}

func App() {
	start := time.Now()
	defer func() { fmt.Printf("Executed time: %s\n", time.Since(start)) }()

	if ok := flag.PraseFlag(); !ok {
		// TODO: error
		os.Exit(1)
	}

	fmt.Printf("Performing donations on %s\n", flag.GetFilePath())

	client, err := payment.GetClient(flag.GetPublickey(), flag.GetSecretkey())

	if err != nil {
		// TODO: error
		os.Exit(1)
	}

	consumers := summaries.GetConsumers(flag.GetNumberTask())

	beginTransaction(client, consumers)

	//summaries
	summaries.PrintSummaries(consumers)

}
