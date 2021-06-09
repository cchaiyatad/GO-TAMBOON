package summaries

import (
	"log"
	S "tamboon/model/summary"
)

func GetConsumers(counts, top int) chan *S.Summary {
	consumers := make(chan *S.Summary, counts)
	for i := 0; i < counts; i++ {
		consumers <- S.CreateNewSummary(top)
	}

	return consumers
}

func GetSummaries(consumers <-chan *S.Summary, top int, counts int, isDebug bool) *S.Summary {
	summaries := S.CreateNewSummary(top)

	// for consumer := range consumers {
	for i := 0; i < counts; i++ {
		consumer := <-consumers
		if isDebug {
			log.Printf("summaries: %s\n", consumer)
		}
		summaries.Merge(consumer)
	}

	return summaries
}
