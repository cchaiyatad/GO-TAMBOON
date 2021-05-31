package decrypt

import (
	"bufio"
	"os"
	"tamboon/cipher"
)

var fp *os.File
var rotData *cipher.Rot128Reader
var producer chan []byte

func Init(filePath string) {

	fp, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	rotData, err = cipher.NewRot128Reader(fp)
	producer = make(chan []byte, 10)
}

//	Format: Name,Amount,Card,CCV,Month,Year
// 	Ex: 	Mr. Bildad R Sackville,5073530,4716972894061735,064,8,2019
func GetDecrypt() {
	scanner := bufio.NewScanner(rotData)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		// fmt.Printf("%s\n", scanner.Text())
		producer <- []byte(scanner.Text())
	}
}

func Producer() <-chan []byte {
	return producer
}

func CloseFile() {
	fp.Close()
}
