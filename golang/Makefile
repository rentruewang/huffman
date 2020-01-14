MAIN = huffman.go

all: default

default: $(MAIN)
	go build

gcc: $(MAIN)
	gccgo -O3 -o golang -static