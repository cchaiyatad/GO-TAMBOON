package payment

// import (
// 	S "tamboon/model/summary"

// 	"github.com/omise/omise-go"
// )

// // var consumers chan *S.Summary

// // For parallel

// func GetPaymentHandler(publicKey, secretKey string, counts int) (*omise.Client, <-chan *S.Summary, error) {
// 	client, err := GetClient(publicKey, secretKey)
// 	if err != nil {
// 		// TODO: Handle Error
// 		return nil, nil, err
// 	}
// 	consumers := makeConsumers(counts)
// 	return client, consumers, nil
// }

// func makeConsumers(counts int) <-chan *S.Summary {
// 	consumers := make(chan *S.Summary, counts)
// 	for i := 0; i < counts; i++ {
// 		consumers <- S.CreateNewSummary()
// 	}

// 	return consumers
// }

// // func chargeParallel(sum *S.Summary, raw []byte) *S.Summary {
// // 	if sum == nil {
// // 		sum, _ = S.CreateNewSummary()
// // 	}

// // 	func(r []byte) {
// // 		tran, err := T.CreateTransaction(r)
// // 		if err == nil {
// // 			isSuccess := charge(tran)
// // 			sum.Update(*tran, isSuccess)
// // 		}
// // 	}(raw)
// // 	return sum
// // }

// // func Run(pd <-chan []byte) {
// // 	for {
// // 		rawData, ok := <-pd
// // 		if !ok || rawData == nil {
// // 			break
// // 		}
// // 		s := <-consumers
// // 		go func(raw []byte) {
// // 			consumers <- chargeParallel(s, raw)
// // 		}(rawData)

// // 	}

// // 	// Summary
// // 	for i := 0; i < cap(consumers); i++ {
// // 		s := <-consumers
// // 		fmt.Printf("%#v\n", s)
// // 	}
// // 	close(consumers)

// // }
