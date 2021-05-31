package payment

import (
	"fmt"
	S "tamboon/model/summary"
	T "tamboon/model/transaction"
)

// var semaphor chan struct{}
var consumers chan *S.Summary

func makeConsumers(counts int) {
	// semaphor = make(chan struct{}, counts)
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

	// semaphor <- struct{}{}
	func(r []byte) {
		tran, err := T.CreateTransaction(r)
		// fmt.Printf("%v:%v:%v:%v\n", string(r), tran, sum, err)
		if err == nil {
			isSuccess := charge(tran)
			sum.Update(*tran, isSuccess)
		}
		// fmt.Println("Here")
	}(raw)
	// <-semaphor
	return sum
}

func Run(pd <-chan []byte) {
	// n := 1
	// for ; n > 0; n-- {
	for {
		rawData, ok := <-pd
		if !ok || rawData == nil {
			break
		}
		s := <-consumers
		// n++
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
