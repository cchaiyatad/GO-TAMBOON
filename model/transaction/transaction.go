package transaction

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
	// check for error
	return nil, nil
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
