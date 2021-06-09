package decrypt

import (
	"bufio"
	"log"
	"os"
	"tamboon/cipher"
)

func GetProducer(filePath string, isDebug bool) (<-chan []byte, *os.File, error) {
	rotData, filePointer, err := getDecryptFile(filePath)
	if err != nil {
		return nil, filePointer, err
	}

	producer := make(chan []byte)
	go beginDecrypt(rotData, producer, isDebug)
	return producer, filePointer, nil
}

func CleanProducer(filePointer *os.File) {
	filePointer.Close()
}

func getDecryptFile(filePath string) (*cipher.Rot128Reader, *os.File, error) {
	filePointer, err := os.Open(filePath)

	if err != nil {
		return nil, nil, err
	}

	rotData, err := cipher.NewRot128Reader(filePointer)

	return rotData, filePointer, err
}

func beginDecrypt(rotData *cipher.Rot128Reader, prod chan<- []byte, isDebug bool) {
	scanner := bufio.NewScanner(rotData)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		txt := []byte(scanner.Text())
		if isDebug {
			log.Printf("decrpyt: %s\n", txt)
		}
		prod <- txt
	}

	// End signal
	prod <- nil
}
