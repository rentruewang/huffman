package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
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

	tree := MakeHuffmanTreeFromFile(infile)
	codes := tree.Huffman()

	var data []byte
	// Write in a json format to the output file
	if data, err = json.Marshal(codes); err != nil {
		// I trust what I did will not break this
		log.Fatalln("unreachable")
	}

	var n int
	if n, err = outfile.Write(data); err != nil {
		log.Fatalf("Read %d lines. Buffer overflow\n", n)
	}
}
