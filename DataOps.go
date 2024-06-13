package utl

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

// Load data from file.
func LoadGameData[T any](httpMethod, apiURL, fileName string) T {
	var file *os.File

	fileInfo := CheckFileExists(fileName)
	defer file.Close()

	if fileInfo != nil && fileInfo.Size() > 0 {

		file = OperateFile(fileName, os.O_RDONLY, 0655)
		return DecodeFileToStruct[T](file)

	} else {

		log.Printf("Creating %s...\n", fileName)
		resp := CallApi(httpMethod, apiURL)
		defer resp.Body.Close()

		file = OperateFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0655)
		_ = WriteToFile(file, resp.Body)

		return DecodeFileToStruct[T](file)
	}
}

// When called takes input and gives a String.
func GetString() func(inpRdr *bufio.Reader) string {
	return func(inpRdr *bufio.Reader) string {
		word, err := inpRdr.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		return strings.ToLower(strings.TrimSpace(word))
	}
}

// When called takes input and gives a Rune.
func GetRune() func(inpRdr *bufio.Reader) rune {
	return func(inpRdr *bufio.Reader) rune {
		r, _, err := inpRdr.ReadRune()
		if err != nil {
			log.Fatalln(err)
		}
		return r
	}
}

// Selects a random item from array
func GetRandItem[T any](tp []T) *T {
	fmt.Println(len(tp))
	c := tp[rand.Intn(len(tp))]
	return &c
}
