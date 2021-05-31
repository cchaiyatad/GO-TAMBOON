package payment

import (
	"log"
	T "tamboon/model/transaction"
	"time"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

var client *omise.Client

func Init(publicKey, secretKey string) {
	var e error
	client, e = omise.NewClient(publicKey, secretKey)
	if e != nil {
		log.Fatal(e)
	}
}

func createToken(tran *T.Transaction) (token *omise.Token, e error) {

	token, createToken := &omise.Token{}, &operations.CreateToken{
		Name:            tran.Name,
		Number:          tran.CardNumber,
		ExpirationMonth: time.Month(tran.Month),
		ExpirationYear:  tran.Year,
		SecurityCode:    tran.CCV,
	}

	if e := client.Do(token, createToken); e != nil {
		log.Fatalln(e)
		return nil, e
	}

	return
}

func Charge(tran *T.Transaction) {

	token, e := createToken(tran)
	if e != nil {
		//TODO: Error
		log.Fatal(e)
	}

	// Creates a charge from the token
	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   int64(tran.Amount),
		Currency: "thb",
		Card:     token.ID,
	}
	if e = client.Do(charge, createCharge); e != nil {
		log.Fatal(e)
	}

	log.Printf("charge: %s  amount: %s %d\n", charge.ID, charge.Currency, charge.Amount)
}
