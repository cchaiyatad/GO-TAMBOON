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
	Donors        D.TopDonors //only top t success
}

func (s *Summary) String() string {
	return fmt.Sprintf("Success: %d\t%d\nFail: %d\t%d\nTop %d Donors: %s",
		s.CountSuccess,
		s.AmountSuccess,
		s.CountFail,
		s.AmountFail,
		len(s.Donors),
		s.Donors)
}

func CreateNewSummary(size int) *Summary {
	donors := make(D.TopDonors, size)
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
		lastIdx := len(s.Donors) - 1
		if s.Donors[lastIdx].Amount < int(t.Amount) {
			s.Donors[lastIdx] = &D.Donor{Name: t.Name, Amount: int(t.Amount)}
			sort.Sort(s.Donors)
		}

	} else {
		s.CountFail += 1
		s.AmountFail += int(t.Amount)
	}
}

func (s1 *Summary) Merge(s2 *Summary) {
	s1.AmountFail += s2.AmountFail
	s1.CountFail += s2.CountFail
	s1.AmountSuccess += s2.AmountSuccess
	s1.CountSuccess += s2.CountSuccess

	topDonor := make(D.TopDonors, len(s1.Donors))
	copy(topDonor, s1.Donors)
	topDonor = append(topDonor, s2.Donors...)
	sort.Sort(topDonor)
	s1.Donors = topDonor[:len(s1.Donors)]
}
