package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
	"strings"
)

// EmptyString is not a string at all
const EmptyString string = ""

func main() {
	var err error

	var infname, outfname string

	flag.StringVar(&infname, "i", EmptyString, "The file to be encoded into huffman coding")
	flag.StringVar(&outfname, "o", EmptyString, "The file where encoded strings can be saved")

	flag.Parse()

	var infile, outfile *os.File

	// Read the input file, then handle the file to MakeHuffmanTreeFromFile
	if infile, err = os.Open(infname); err != nil {
		log.Fatalln("No such input file!")
	}
	defer infile.Close()

	// Output is a json file.
	if outfile, err = os.Create(outfname); err != nil {
		log.Fatalln("No such output file!")
	}
	defer outfile.Close()

	var data []byte

	// As of Go 1.16, the function stays in io package
	data, err = io.ReadAll(infile)
	if err != nil {
		log.Fatalln("There is something terrifying in the input file!")
	}

	// Still going to do some processing before using
	tmpData := string(data)

	// Newlines are not very printable, so I'll just get rid of it.
	content := strings.ReplaceAll(tmpData, "\n", " ")

	// This does all the heavy lifting.
	codes := Huffman(content)

	// Write in a json format to the output file
	if data, err = json.Marshal(codes); err != nil {
		// Come on! It's just a map[string]string to a json string
		// What could possibly go wrong?
		log.Fatalln("This is not possible!")
	}

	var n int
	if n, err = outfile.Write(data); err != nil {
		log.Fatalf("Read %d lines. Buffer overflow!\n", n)
	}
}
