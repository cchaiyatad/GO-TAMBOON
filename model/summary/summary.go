package summary

import (
	"fmt"
	"sort"
	D "tamboon/model/donor"
	T "tamboon/model/transaction"
)

type Summary struct {
	countSuccess  int
	amountSuccess int
	countFail     int
	amountFail    int
	donors        D.TopDonors //only top t success
}

func (s *Summary) String() string {
	return fmt.Sprintf("Success: %d\t%d\nFail: %d\t%d\nTop %d Donors: %s",
		s.countSuccess,
		s.amountSuccess,
		s.countFail,
		s.amountFail,
		len(s.donors),
		s.donors)
}

func CreateNewSummary(size int) *Summary {
	donors := make(D.TopDonors, size)
	for i := range donors {
		donors[i] = &D.Donor{}
	}

	return &Summary{
		donors: donors,
	}
}

func (s *Summary) Update(t T.Transaction, isSuccess bool) {
	if isSuccess {
		s.countSuccess += 1
		s.amountSuccess += int(t.Amount)

		// Check max
		lastIdx := len(s.donors) - 1
		if s.donors[lastIdx].Amount < int(t.Amount) {
			s.donors[lastIdx] = &D.Donor{Name: t.Name, Amount: int(t.Amount)}
			sort.Sort(s.donors)
		}

	} else {
		s.countFail += 1
		s.amountFail += int(t.Amount)
	}
}

func (s1 *Summary) Merge(s2 *Summary) {
	s1.amountFail += s2.amountFail
	s1.countFail += s2.countFail
	s1.amountSuccess += s2.amountSuccess
	s1.countSuccess += s2.countSuccess

	topDonor := make(D.TopDonors, len(s1.donors))
	copy(topDonor, s1.donors)
	topDonor = append(topDonor, s2.donors...)
	sort.Sort(topDonor)
	s1.donors = topDonor[:len(s1.donors)]
}

func convertAmountFormat(amount int) float64 {
	return float64(amount) / 100.0
}

func (s *Summary) GetTotalReceived() float64 {
	return convertAmountFormat(s.amountSuccess + s.amountFail)
}

func (s *Summary) GetAmountSuccess() float64 {
	return convertAmountFormat(s.amountSuccess)
}

func (s *Summary) GetAmountFail() float64 {
	return convertAmountFormat(s.amountFail)
}

func (s *Summary) GetTotalCount() int {
	return s.countSuccess + s.countFail
}

func (s *Summary) GetSuccessRate() float64 {
	return float64(s.countSuccess) / float64(s.GetTotalCount()) * 100
}

func (s *Summary) GetAvg() float64 {
	return s.GetAmountSuccess() / float64(s.GetTotalCount())
}
