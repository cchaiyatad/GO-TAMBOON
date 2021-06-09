package controller

import (
	"fmt"
	"log"
	"tamboon/model/summary"
	"tamboon/model/transaction"
	"tamboon/service/decrypt"
	"tamboon/service/flag"
	"tamboon/service/payment"
	"tamboon/service/summaries"
	"time"

	"github.com/omise/omise-go"
)

func beginTransaction(client *omise.Client, consumers chan *summary.Summary, isDebug bool) {
	producer, filePointer, err := decrypt.GetProducer(flag.GetFilePath(), isDebug)
	if err != nil {
		decrypt.CleanProducer(filePointer)
		log.Fatalln(err)
	}

	for line := range producer {
		if isDebug {
			log.Printf("app(line): %s\n", line)
		}

		if line == nil {
			// EOF
			break
		}

		// not in a format Ex. header, expired card
		tran, err := transaction.CreateTransaction(line)
		if err != nil {
			log.Println(err)
		}

		// only continue here for invalid format like header
		if tran == nil {
			continue
		}

		if isDebug && tran != nil {
			log.Printf("app(transaction): %s\n", tran)
		}

		var payErr error

		if err == nil {
			payErr = payment.BeginCharge(tran, client)
		}

		if payErr != nil {
			log.Printf("payment error: %s\n", payErr)
		}

		consumer := <-consumers
		consumer.Update(*tran, payErr == nil && err != nil)
		consumers <- consumer
	}
	fmt.Printf("Done.\n\n")
	decrypt.CleanProducer(filePointer)
	summaries.CleanConsumer(consumers)
}

func App() {
	start := time.Now()
	defer func() { fmt.Printf("\nExecuted time: %s\n", time.Since(start)) }()

	if err := flag.PraseFlag(); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Performing donations on %s\n", flag.GetFilePath())

	client, err := payment.GetClient(flag.GetPublickey(), flag.GetSecretkey())

	if err != nil {
		log.Fatalln(err)
	}

	consumers := summaries.GetConsumers(flag.GetNumberTask(), flag.GetTopsNumber())

	beginTransaction(client, consumers, flag.IsDebug())

	sum := summaries.GetSummaries(consumers, flag.GetTopsNumber(), flag.IsDebug())
	sum.PrintSummaries()
}
