package transaction

import (
	"bytes"
	"os"
	"strconv"
)

// 	Ex: 	Mr. Bildad R Sackville,5073530,4716972894061735,064,8,2019
type Transaction struct {
	Name       string
	Amount     int64
	CardNumber string
	CCV        string
	Month      int
	Year       int
}

func CreateTransaction(d []byte) (*Transaction, error) {
	field := bytes.Split(d, []byte(","))
	// check for error

	amount, err := strconv.Atoi(string(field[1]))
	if err != nil {
		os.Exit(1)
	}
	month, err := strconv.Atoi(string(field[4]))
	if err != nil {
		os.Exit(1)
	}

	year, err := strconv.Atoi(string(field[5]))
	if err != nil {
		os.Exit(1)
	}

	return &Transaction{
		Name:       string(field[0]),
		Amount:     int64(amount),
		CardNumber: string(field[2]),
		CCV:        string(field[3]),
		Month:      month,
		Year:       year,
	}, nil

}

func CreateTestTransaction() (*Transaction, error) {
	// check for error
	return &Transaction{
		Name:       "Ms. Semolina B Brockhouse",
		Amount:     2674018,
		CardNumber: "4539823809118158",
		CCV:        "260",
		Month:      10,
		Year:       2022,
	}, nil
}
