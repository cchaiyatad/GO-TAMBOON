package controller

import (
	"fmt"
	"os"
	"tamboon/service/decrypt"
	"tamboon/service/flag"
	"time"
)

func App() {
	start := time.Now()
	defer func() { fmt.Printf("Executed time: %s\n", time.Since(start)) }()

	ok := flag.PraseFlag()

	if !ok {
		os.Exit(1)
	}
	fmt.Printf("Performing donations on %s\n", flag.GetFilePath())

	decrypt.Init(flag.GetFilePath())
	defer decrypt.CloseFile()

	decrypt.GetDecrypt()
}
