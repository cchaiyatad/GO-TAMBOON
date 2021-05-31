package decrypt

import (
	"bufio"
	"fmt"
	"os"
	"tamboon/cipher"
)

var fp *os.File
var rotData *cipher.Rot128Reader

func Init(filePath string) {

	fp, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	rotData, err = cipher.NewRot128Reader(fp)

}

//	Format: Name,Amount,Card,CCV,Month,Year
// 	Ex: 	Mr. Bildad R Sackville,5073530,4716972894061735,064,8,2019
func GetDecrypt() {
	scanner := bufio.NewScanner(rotData)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
}

func CloseFile() {
	fp.Close()
}
