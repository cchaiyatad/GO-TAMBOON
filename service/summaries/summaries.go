package summaries

import (
	"fmt"
	S "tamboon/model/summary"
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
	for consumer := range consumers {
		fmt.Printf("%s\n", consumer)
	}
}
