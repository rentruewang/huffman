package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var err error

	var inputName, outputName string

	flag.StringVar(&inputName, "i", "", "The file to be encoded into huffman coding")
	flag.StringVar(&outputName, "o", "", "The file where encoded strings can be saved")

	flag.Parse()

	var inputFile, outputFile *os.File

	if inputFile, err = os.Open(inputName); err != nil {
		log.Fatalln("No such input file!")
	}
	defer inputFile.Close()

	if outputFile, err = os.Create(outputName); err != nil {
		log.Fatalln("No such output file!")
	}
	defer outputFile.Close()

	var data []byte

	// As of Go 1.16, the function stays in io package
	data, err = io.ReadAll(inputFile)
	if err != nil {
		log.Fatalln("There is something terrifying in the input file!")
	}

	tmpData := string(data)

	// Newlines are not very printable, so don't care about it.
	content := strings.ReplaceAll(tmpData, "\n", " ")

	// Compute Huffman codes.
	codes := Huffman(content)

	if data, err = json.Marshal(codes); err != nil {
		log.Fatalln("map[string]string to json fails!")
	}

	var n int

	if n, err = outputFile.Write(data); err != nil {
		log.Fatalf("Read %d lines. Buffer overflow!\n", n)
	}
}
