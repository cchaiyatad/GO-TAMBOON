package summary

import T "tamboon/model/transaction"

type Summary struct {
	CountSuccess  int
	AmountSuccess int
	CountFail     int
	AmountFail    int
	MaxAmount     int //only success
	MaxName       string
}

func CreateNewSummary() *Summary {
	return &Summary{
		CountSuccess:  0,
		AmountSuccess: 0,
		CountFail:     0,
		AmountFail:    0,
		MaxAmount:     0,
		MaxName:       "",
	}
}

func (s *Summary) Update(t T.Transaction, isSuccess bool) {
	if isSuccess {
		s.CountSuccess += 1
		s.AmountSuccess += int(t.Amount)

		// Check max
		if s.MaxAmount < int(t.Amount) {
			s.MaxAmount = int(t.Amount)
			s.MaxName = t.Name
		}
	} else {
		s.CountFail += 1
		s.AmountFail += int(t.Amount)
	}
}
