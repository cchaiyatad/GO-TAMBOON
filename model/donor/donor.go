package donor

import "fmt"

type Donor struct {
	Name   string
	Amount int
}

func (donor *Donor) String() string {
	return fmt.Sprintf("%s: %d", donor.Name, donor.Amount)
}

type TopDonors []*Donor

func (donors TopDonors) Len() int           { return len(donors) }
func (donors TopDonors) Swap(i, j int)      { donors[i], donors[j] = donors[j], donors[i] }
func (donors TopDonors) Less(i, j int) bool { return donors[i].Amount > donors[j].Amount }
