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

Top 3 Donors: 
1.%s 		THB %.2f
2.%s 		THB %.2f
3.%s 		THB %.2f
`
)

func GetConsumers(counts int) chan *S.Summary {
	consumers := make(chan *S.Summary, counts)
	for i := 0; i < counts; i++ {
		consumers <- S.CreateNewSummary()
	}

	return consumers
}

func CleanConsumer(consumers chan *S.Summary) {
	close(consumers)
}

func PrintSummaries(consumers <-chan *S.Summary) {
	summaries := S.CreateNewSummary()

	for consumer := range consumers {
		// fmt.Printf("%s\n", consumer)
		summaries.Merge(consumer)
	}
	fmt.Printf(templateFormat,
		(float64(summaries.AmountSuccess+summaries.AmountFail) / 100.0),
		(float64(summaries.AmountSuccess) / 100.0),
		(float64(summaries.AmountFail) / 100.0),
		(summaries.CountSuccess + summaries.CountFail),
		(float64(summaries.CountSuccess) / float64(summaries.CountSuccess+summaries.CountFail) * 100.0),
		(float64(summaries.AmountSuccess+summaries.AmountFail) / float64(summaries.CountSuccess+summaries.CountFail)),
		summaries.Donors[0].Name,
		(float64(summaries.Donors[0].Amount) / 100.0),
		summaries.Donors[1].Name,
		(float64(summaries.Donors[1].Amount) / 100.0),
		summaries.Donors[2].Name,
		(float64(summaries.Donors[2].Amount) / 100.0))
}
