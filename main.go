package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("message.txt")
	if err != nil {
		// if error is present
		log.Fatal("error", err)
	}

	str := ""

	for { // while the file is being read, we will make a buffer (data) of byte[], size 8
		data := make([]byte, 8)
		n, err := file.Read(data) // every time we read, we update n and err
		if err != nil {
			break
		}

		data = data[:n]
		if i := bytes.IndexByte(data, '\n'); i != -1 { // while the position i in data is not EOF
			str += string(data[:i]) // append the str with the data up from 0 -> i == get the whole line
			data = data[i+1:] // get the next line
			fmt.Printf("read: %s\n", str) 
			str = ""
		}

		str += string(data)
	}

	if len(str) != 0 { // if there's something left
		fmt.Printf("read: %s\n", str)
	}
}