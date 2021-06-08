package decrypt

import (
	"bufio"
	"fmt"
	"os"
	"tamboon/cipher"
)

func GetProducer(filePath string) (<-chan []byte, *os.File) {
	rot, fp, _ := getDecryptFile(filePath)
	// TODO: Check error

	prod := make(chan []byte)
	go beginDecrypt(rot, prod)
	return prod, fp
}

func CleanProducer(fp *os.File) {
	fp.Close()
}

func getDecryptFile(filePath string) (*cipher.Rot128Reader, *os.File, error) {
	fp, err := os.Open(filePath)

	if err != nil {
		// TODO: handle error
		return nil, nil, err
	}

	rot, err := cipher.NewRot128Reader(fp)

	// TODO: handle error
	return rot, fp, err
}

func beginDecrypt(rot *cipher.Rot128Reader, prod chan<- []byte) {
	scanner := bufio.NewScanner(rot)
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
