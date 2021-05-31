package flag

import (
	"flag"
	"fmt"
)

var file = flag.String("f", "", "path to file")
var pk = flag.String("pk", "", "omise Public Key")
var sk = flag.String("sk", "", "omise Secret Key")
var number = flag.Int("n", 8, "number of tasks")

func PraseFlag() (ok bool) {
	flag.Parse()

	if *file != "" && *pk != "" && *sk != "" {
		ok = true
	} else {
		// throw error
		fmt.Println("Usage: ./app -f [pathToFile] -pk [omisePublicKey] -sk [omiseSecretKey]")
	}

	return
}

func GetFilePath() string {
	return *file
}

func GetNumberTask() int {
	return *number
}
func GetPublickey() string {
	return *pk
}
func GetSecretkey() string {
	return *sk
}
