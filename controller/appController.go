package controller

import (
	"fmt"
	"log"
	S "tamboon/model/summary"
	"tamboon/model/transaction"
	T "tamboon/model/transaction"
	"tamboon/service/decrypt"
	"tamboon/service/flag"
	"tamboon/service/payment"
	"tamboon/service/summaries"
	"time"

	"github.com/omise/omise-go"
)

func beginTransaction(client *omise.Client, consumers chan *S.Summary, isDebug bool) {
	producer, filePointer, err := decrypt.GetProducer(flag.GetFilePath(), isDebug)
	if err != nil {
		decrypt.CleanProducer(filePointer)
		log.Fatalln(err)
	}

	for line := range producer {
		if isDebug {
			log.Printf("app(line): %s\n", line)
		}

		// EOF
		if line == nil {
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

		consumer := <-consumers
		go consume(client, consumers, consumer, tran, err == nil, isDebug)

	}
	fmt.Printf("Done.\n\n")
	decrypt.CleanProducer(filePointer)
}

func consume(client *omise.Client,
	consumers chan *S.Summary,
	consumer *S.Summary,
	tran *T.Transaction,
	doCharge bool,
	isDebug bool) {

	var payErr error

	if doCharge {
		payErr = payment.BeginCharge(tran, client)
	}

	if payErr != nil {
		log.Printf("payment error: %s\n", payErr)
	}

	consumer.Update(tran, payErr == nil && doCharge, isDebug)
	consumers <- consumer
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

	sum := summaries.GetSummaries(consumers, flag.GetTopsNumber(), flag.GetNumberTask(), flag.IsDebug())
	sum.PrintSummaries()
}
