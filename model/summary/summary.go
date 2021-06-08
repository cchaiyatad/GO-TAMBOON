package summary

import (
	"fmt"
	"sort"
	D "tamboon/model/donor"
	T "tamboon/model/transaction"
)

type Summary struct {
	CountSuccess  int
	AmountSuccess int
	CountFail     int
	AmountFail    int
	Donors        D.TopDonors //only top 3 success
}

func (s *Summary) String() string {
	return fmt.Sprintf("Success: %d\t%d\nFail: %d\t%d\nTop 3 Donors: %s", s.CountSuccess, s.AmountSuccess, s.CountFail, s.AmountFail, s.Donors)
}

func CreateNewSummary() *Summary {
	donors := make(D.TopDonors, 3)
	for i := range donors {
		donors[i] = &D.Donor{}
	}

	return &Summary{
		Donors: donors,
	}
}

func (s *Summary) Update(t T.Transaction, isSuccess bool) {
	if isSuccess {
		s.CountSuccess += 1
		s.AmountSuccess += int(t.Amount)

		// Check max
		if s.Donors[2].Amount < int(t.Amount) {
			s.Donors[2] = &D.Donor{Name: t.Name, Amount: int(t.Amount)}
			sort.Sort(s.Donors)
		}

	} else {
		s.CountFail += 1
		s.AmountFail += int(t.Amount)
	}
}
