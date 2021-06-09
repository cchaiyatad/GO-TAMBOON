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

func CleanConsumer(consumers chan *S.Summary) {
	close(consumers)
}

func GetSummaries(consumers <-chan *S.Summary, top int, isDebug bool) *S.Summary {
	summaries := S.CreateNewSummary(top)

	for consumer := range consumers {
		if isDebug {
			log.Printf("summaries: %s\n", consumer)
		}
		summaries.Merge(consumer)
	}
	return summaries
}
