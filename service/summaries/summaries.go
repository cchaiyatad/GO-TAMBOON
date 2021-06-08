package summaries

import (
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
