package goutility

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

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

func GetNewRdr(reader *bufio.Reader) *bufio.Reader {
	return bufio.NewReader(reader)
}

func GetNewRdrWritr(reader *bufio.Reader, writer *bufio.Writer) *bufio.ReadWriter {
	return bufio.NewReadWriter(reader, writer)
}

func GetNewStdInRdr() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}
