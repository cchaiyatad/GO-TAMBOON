package payment

import (
	"fmt"
	S "tamboon/model/summary"
	T "tamboon/model/transaction"
)

var consumers chan *S.Summary

func makeConsumers(counts int) {
	consumers = make(chan *S.Summary, counts)
	for i := 0; i < counts; i++ {
		go func() {
			s, _ := S.CreateNewSummary()
			consumers <- s
		}()
	}
}

func chargeParallel(sum *S.Summary, raw []byte) *S.Summary {
	if sum == nil {
		sum, _ = S.CreateNewSummary()
	}

	func(r []byte) {
		tran, err := T.CreateTransaction(r)
		if err == nil {
			isSuccess := charge(tran)
			sum.Update(*tran, isSuccess)
		}
	}(raw)
	return sum
}

func Run(pd <-chan []byte) {
	for {
		rawData, ok := <-pd
		if !ok || rawData == nil {
			break
		}
		s := <-consumers
		go func(raw []byte) {
			consumers <- chargeParallel(s, raw)
		}(rawData)

	}

	// Summary
	for i := 0; i < cap(consumers); i++ {
		s := <-consumers
		fmt.Printf("%#v\n", s)
	}
	close(consumers)

}
