package payment

import (
	T "tamboon/model/transaction"
	"time"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

func GetClient(publicKey, secretKey string) (*omise.Client, error) {
	return omise.NewClient(publicKey, secretKey)
}

func BeginCharge(raw []byte, client *omise.Client) error {
	tran, err := T.CreateTransaction(raw)

	if err != nil {
		return err
	}

	token, err := createToken(tran, client)
	if err != nil {
		return err
	}

	return doCharge(token, tran, client)
}

func createToken(tran *T.Transaction, client *omise.Client) (*omise.Token, error) {
	token, createToken := &omise.Token{}, &operations.CreateToken{
		Name:            tran.Name,
		Number:          tran.CardNumber,
		ExpirationMonth: time.Month(tran.Month),
		ExpirationYear:  tran.Year,
		SecurityCode:    tran.CCV,
	}

	if err := client.Do(token, createToken); err != nil {
		return nil, err
	}

	return token, nil
}

func doCharge(token *omise.Token, tran *T.Transaction, client *omise.Client) error {
	// Creates a charge from the token
	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   int64(tran.Amount),
		Currency: "thb",
		Card:     token.ID,
	}

	// if err == mil mean finish
	return client.Do(charge, createCharge)
}
