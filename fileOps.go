package utl

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
)

// Opens a file and makes it available in byte array
func LoadFile(flPth string) []byte {
	flBf, err := os.ReadFile(flPth)
	if err != nil {
		log.Fatalf("Failed to resolve config path: %v", err)
		return nil
	} else {
		return flBf
	}
}

// Open, Create, Read, Append
func OperateFile(fileName string, openPerms int, filePerms os.FileMode) *os.File {
	file, err := os.OpenFile(fileName, openPerms, filePerms)
	if err != nil {
		log.Fatalln(err)
	}
	return file
}

func WriteToFile(file *os.File, resp io.Reader) int {
	defer file.Close()
	var writenSize int
	var err error

	responseScanner := bufio.NewScanner(resp)
	for responseScanner.Scan() {
		writenSize, err = file.Write(responseScanner.Bytes())
		if err != nil {
			log.Fatalln(err)
			return 0
		}
	}
	return writenSize
}

func DecodeFileToStruct[T any](file *os.File) T {
	defer file.Close()
	var result T
	decodedJson := json.NewDecoder(bufio.NewReader(file))
	for decodedJson.More() {
		err := decodedJson.Decode(&result)
		if err != nil {
			return result
		}
	}
	return result
}

// Check if file already exists
func CheckFileExists(fileName string) os.FileInfo {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Printf("File Not found %v\n", err)
		return nil
	} else {
		log.Printf("Loading %s...\n", fileName)
		return fileInfo
	}
}
