package flag

import (
	"flag"
	"fmt"
)

var file = flag.String("f", "", "path to file")

func PraseFlag() (ok bool) {
	flag.Parse()

	if *file != "" {
		ok = true
	} else {
		// throw error
		fmt.Println("Usage: ./app -f [pathToFile]")
	}

	return
}

func GetFilePath() string {
	return *file
}
