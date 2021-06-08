package decrypt

import (
	"bufio"
	"fmt"
	"os"
	"tamboon/cipher"
)

func GetProducer(filePath string) (<-chan []byte, *os.File, error) {
	rotData, filePointer, err := getDecryptFile(filePath)
	if err != nil {
		return nil, filePointer, err
	}

	producer := make(chan []byte)
	go beginDecrypt(rotData, producer)
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

func beginDecrypt(rotData *cipher.Rot128Reader, prod chan<- []byte) {
	scanner := bufio.NewScanner(rotData)
	scanner.Split(bufio.ScanLines)
	fmt.Printf("%s\n", "Hi")
	for scanner.Scan() {
		txt := []byte(scanner.Text())
		// fmt.Printf("%s\n", txt)
		prod <- txt
	}

	// End signal
	prod <- nil
}
