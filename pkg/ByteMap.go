package pkg

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Hideous code, but it runs. Created for a how-to YouTube video.

func EncodePayload(key string, payloadname string, payload string) {
	// Define Hex location map variable
	hexLocation := make(map[byte]string)

	// Read key file contents
	keyFile, err := ioutil.ReadFile(key)
	if err != nil {
		log.Fatal(err)
	}

	// Build hexLocation map for the key file
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

	// Check if all 256 bytes are present in the map, if not error out and close
	if len(hexLocation) < 256 {
		// I can implement a check to see which bytes, but I am lazy.
		fmt.Println("[*] Cound not complete location hex map, some bytes are misisng from the key file!")
		os.Exit(1)
	}

	payloadFile, err := ioutil.ReadFile(payload)
	if len(payloadFile) == 0 {
		// If not a file, try to load the string as hex payload
		payloadFile, err = hex.DecodeString(payload)
		if err != nil {
			fmt.Printf("Error decoding arg 1: %s\n", err)
			os.Exit(1)
		}
	}

	// Open the encoded payload file on disk
	file, err := os.OpenFile(payloadname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("Failed to create payload file")
	}

	// Create writer interface
	writer := bufio.NewWriter(file)

	// For every byte in payload file, write the location of the byte in the key file
	for _, fileByte := range payloadFile {
		_, _ = writer.WriteString(hexLocation[fileByte] + ",")
	}

	writer.Flush()
	file.Close()
}

func DecodePayload(key string, encodedPayload string) bytes.Buffer {

	// Read key file contents
	keyFile, err := ioutil.ReadFile(key)
	if err != nil {
		log.Fatal(err)
	}

	// Read encoded payload file contents
	encodedPayloadFile, err := ioutil.ReadFile(encodedPayload)
	if err != nil {
		log.Fatal(err)
	}

	// Remove the last "," and split the data by ","
	textData := strings.Split(strings.TrimSuffix(string(encodedPayloadFile), ","), ",")

	var shellcode bytes.Buffer

	// Decode the shellcode
	for _, fileByte := range textData {
		i, _ := strconv.Atoi(fileByte)
		shellcode.WriteByte(keyFile[i])
	}

	return shellcode

}
