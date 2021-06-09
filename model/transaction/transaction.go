package transaction

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

type Transaction struct {
	Name       string
	Amount     int64
	CardNumber string
	CCV        string
	Month      int
	Year       int
}

func (t *Transaction) String() string {
	return fmt.Sprintf("transaction: name:%s amount:%.2f cardNo:%s ccv:%s expire:%d/%d",
		t.Name,
		float64(t.Amount)/100,
		t.CardNumber,
		t.CCV,
		t.Month,
		t.Year,
	)
}

func parseNumber(text, field string) (int, error) {
	num, err := strconv.Atoi(text)
	if err != nil {
		return 0, fmt.Errorf("create transaction error: %s must be number", field)
	}
	return num, nil
}

// 	Ex: Mr. Bildad R Sackville,5073530,4716972894061735,064,8,2019
func CreateTransaction(d []byte) (*Transaction, error) {
	field := bytes.Split(d, []byte(","))

	// check for error
	amount, err := parseNumber(string(field[1]), "Amount")
	if err != nil {
		return nil, err
	}

	month, err := parseNumber(string(field[4]), "Month")
	if err != nil {
		return nil, err
	}

	year, err := parseNumber(string(field[5]), "Year")
	if err != nil {
		return nil, err
	}

	// Check expired date
	var expiredError error

	tranDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	if tranDate.Before(time.Now()) {
		expiredError = fmt.Errorf("create transaction error: expiration date cannot be in the past")
	}

	return &Transaction{
		Name:       string(field[0]),
		Amount:     int64(amount),
		CardNumber: string(field[2]),
		CCV:        string(field[3]),
		Month:      month,
		Year:       year,
	}, expiredError
}
