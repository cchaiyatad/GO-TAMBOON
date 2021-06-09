package summaries

import (
	"fmt"
	S "tamboon/model/summary"
)

const (
	templateFormat = `Summary:
total received:		THB %.2f
successfully donated:	THB %.2f
faulty donation:	THB %.2f

all:			%d times
success rate:		%.2f%%
average per person:	THB %.2f

`
	topDonorHeaderFormat = "Top %d Donors:\n"
	topDonorFormat       = "%d. %s 		THB %.2f\n"
)

func GetConsumers(counts, top int) chan *S.Summary {
	consumers := make(chan *S.Summary, counts)
	for i := 0; i < counts; i++ {
		consumers <- S.CreateNewSummary(top)
	}

	return consumers
}

func CleanConsumer(consumers chan *S.Summary) {
	close(consumers)
}

func GetSummaries(consumers <-chan *S.Summary, top int) *S.Summary {
	summaries := S.CreateNewSummary(top)

	for consumer := range consumers {
		// fmt.Printf("%s\n", consumer)
		summaries.Merge(consumer)
	}
	return summaries
}

func PrintSummaries(summaries *S.Summary) {
	printStatus(summaries)
	printTopDonor(summaries)
}

func printStatus(summaries *S.Summary) {
	fmt.Printf(templateFormat,
		(float64(summaries.AmountSuccess+summaries.AmountFail) / 100.0),
		(float64(summaries.AmountSuccess) / 100.0),
		(float64(summaries.AmountFail) / 100.0),
		(summaries.CountSuccess + summaries.CountFail),
		(float64(summaries.CountSuccess) / float64(summaries.CountSuccess+summaries.CountFail) * 100.0),
		(float64(summaries.AmountSuccess+summaries.AmountFail) / float64(summaries.CountSuccess+summaries.CountFail)))

}

func printTopDonor(summaries *S.Summary) {
	fmt.Printf(topDonorHeaderFormat, len(summaries.Donors))
	for i := range summaries.Donors {
		fmt.Printf(topDonorFormat,
			i,
			summaries.Donors[i].Name,
			(float64(summaries.Donors[0].Amount) / 100.0))
	}
}
