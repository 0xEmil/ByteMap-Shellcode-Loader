package pkg

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type LBD struct {
}

func (s *LBD) EncodePayload(key string, payload string) {
	//Define Hex location map variable
	hexLocation := make(map[byte]string)

	//Read key file contents
	keyFile, err := ioutil.ReadFile(key)
	if err != nil {
		log.Fatal(err)
	}

	//Build hexLocation map for the key file
	for counter, fileByte := range keyFile {
		if len(hexLocation) == 256 {
			fmt.Println("[*] Done building location hex map.")
			break
		} else {
			_, ok := hexLocation[byte(fileByte)]
			if !ok {
				hexLocation[byte(fileByte)] = fmt.Sprint(counter)
			}
		}
	}

	//Check if all 256 bytes are present in the map, if not error out and close
	if len(hexLocation) < 256 {
		fmt.Println("[*] Cound not complete location hex map, some bytes are misisng from the key file!")
		os.Exit(1)
	}

	//Read payload file contents
	payloadFile, err := ioutil.ReadFile(payload)
	if err != nil {
		log.Fatal(err)
	}

	//Open the payload.txt file on disk
	file, err := os.OpenFile("payload.encoded", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("Failed to create payload file")
	}

	//Create writer interface
	writer := bufio.NewWriter(file)

	//For byte in payload file, write the location of the byte in the key file
	for _, fileByte := range payloadFile {
		_, _ = writer.WriteString(hexLocation[fileByte] + ",")
	}

	writer.Flush()
	file.Close()
}

func (s *LBD) DecodePayload(key string, encodedPayload string) {

	//Read key file contents
	keyFile, err := ioutil.ReadFile(key)
	if err != nil {
		log.Fatal(err)
	}

	//Read encoded payload file contents
	encodedPayloadFile, err := ioutil.ReadFile(encodedPayload)
	if err != nil {
		log.Fatal(err)
	}

	//Remove the last "," and split the data by ","
	textData := strings.Split(strings.TrimSuffix(string(encodedPayloadFile), ","), ",")

	//Open handler to payload.exe file on disk.
	file, err := os.OpenFile("payload.exe", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("[*] Failed to create payload file.")
		os.Exit(1)
	}

	//Create writer interface
	writer := bufio.NewWriter(file)

	//Write the decoded payload to disk
	for _, fileByte := range textData {
		i, _ := strconv.Atoi(fileByte)
		writer.WriteByte(keyFile[i])
	}

	writer.Flush()
	file.Close()
}
