package goutility

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
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

// Open, Create, Read, Append on a File
func OperateFile(fileName string, openPerms int, filePerms os.FileMode) *os.File {
	file, err := os.OpenFile(fileName, openPerms, filePerms)
	if err != nil {
		log.Fatalln(err)
	}
	return file
}

// Write the data to file.
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

// Load data from file.
func CreateOrLoadData[T any](httpMethod, apiURL, fileName string) T {
	var file *os.File

	fileInfo := CheckFileExists(fileName)
	defer file.Close()

	if fileInfo != nil && fileInfo.Size() > 0 {

		file = OperateFile(fileName, os.O_RDONLY, 0655)
		return RdJsonFileToStruct[T](file)

	} else {

		log.Printf("Creating %s...\n", fileName)
		resp := CallApi(httpMethod, apiURL)
		defer resp.Body.Close()

		file = OperateFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0655)
		_ = WriteToFile(file, resp.Body)

		return RdJsonFileToStruct[T](file)
	}
}

/*
	#### E N C O D E ####   *	#### D E C O D E ####
*/

/* Y  A  M  L */

// Yaml file to struct
func RdYamlFileToStruct[T any](file *os.File) T {
	defer file.Close()
	var result T
	decodedYaml := yaml.NewDecoder(bufio.NewReader(file))
	err := decodedYaml.Decode(&result)
	if err != nil {
		return result
	}
	return result
}

// Yaml buffer to struct
func RdYamlToStruct[T any](Buf []byte) *T {
	var ymlStr T
	err := yaml.Unmarshal(Buf, &ymlStr)
	if err != nil {
		log.Fatalln(err)
	}
	return &ymlStr
}

// Struct to Yaml File
func WrtStructToYamlFile[T any](CustStruct T, file *os.File) {
	defer file.Close()
	encodeYaml := yaml.NewEncoder(bufio.NewWriter(file))
	defer encodeYaml.Close()
	err := encodeYaml.Encode(CustStruct)
	if err != nil {
		fmt.Println("File Written!!")
	} else {
		log.Fatalln("Unable to Write!!", err)
	}
}

// Struct to Yaml buffer
func WrtStructToYaml[T any](ymlStr T) []byte {
	bArr, err := yaml.Marshal(ymlStr)
	if err != nil {
		log.Fatalln(err)
	}
	return bArr
}

/* J  S  O  N */

// Json file to struct
func RdJsonFileToStruct[T any](file *os.File) T {
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

// Json buffer to struct
func RdJsonToStruct[T any](Buf []byte) *T {
	var jsnStr T
	err := json.Unmarshal(Buf, &jsnStr)
	if err != nil {
		log.Fatalln(err)
	}
	return &jsnStr
}

// Struct to Json file
func WrtStructToJsonFile[T any](CustStruct T, file *os.File) {
	defer file.Close()
	encodedJson := json.NewEncoder(bufio.NewWriter(file))
	err := encodedJson.Encode(&CustStruct)
	if err != nil {
		fmt.Println("File Written!!")
	} else {
		log.Fatalln("Unable to Write!!", err)
	}
}

// Struct to Json buffer
func WrtStructToJson[T any](Buf T) []byte {
	bArr, err := json.Marshal(&Buf)
	// err := json.Unmarshal(Buf, &jsnStr)
	if err != nil {
		log.Fatalln(err)
	}
	return bArr
}
