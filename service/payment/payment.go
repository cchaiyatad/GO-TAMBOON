package payment

import (
	"log"
	T "tamboon/model/transaction"
	"time"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

var client *omise.Client

func Init(publicKey, secretKey string, counts int) {
	var e error
	client, e = omise.NewClient(publicKey, secretKey)
	if e != nil {
		log.Fatal(e)
	}
	makeConsumers(counts)
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
		// log.Fatalln(e)
		return nil, e
	}

	return
}

func charge(tran *T.Transaction) bool {

	token, e := createToken(tran)

	if e != nil {
		return false
		//TODO: Error
		// fmt.Println(tran)
		// log.Fatal(e)
	}

	// Creates a charge from the token
	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   int64(tran.Amount),
		Currency: "thb",
		Card:     token.ID,
	}
	e = client.Do(charge, createCharge)
	// if e != nil {
	// 	log.Fatal(e)
	// }
	return e == nil
}
