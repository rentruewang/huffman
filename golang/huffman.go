package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	defer file.Close()
	if err != nil {
		panic("FileNotFound")
	}
	stream, err := ioutil.ReadAll(file)
	str := strings.NewReplacer("\n", "").Replace(string(stream))

	wordCount := make(map[string]int)
	for _, char := range str {
		wordCount[string(char)]++
	}
	fmt.Println(wordCount)

	queue := make([]*node, 0)
	for key, val := range wordCount {
		queue = append(queue, &node{char: key, count: val})
	}
	for i := (len(queue) / 2) - 1; i >= 0; i-- {
		downHeap(queue, i)
	}
	for i, total := 0, len(queue)-1; i < total; i++ {
		var first, second *node
		queue, first = pop(queue)
		queue, second = pop(queue)
		n := &node{left: first, right: second, char: "", count: first.count + second.count}
		idx := len(queue)
		queue = append(queue, n)
		for queue[idx].count < queue[idx/2].count {
			queue[idx], queue[idx/2] = queue[idx/2], queue[idx]
			idx = idx / 2
		}
	}
	huffman := walk(queue[0])
	fmt.Println(huffman)

	total := 0
	for key, str := range huffman {
		total += len(str) * wordCount[key]
	}
	fmt.Println("cost", total)
}

func walkRecursive(tree *node, huffman map[string]string, path string) {
	if tree.char == "" {
		walkRecursive(tree.left, huffman, path+"0")
		walkRecursive(tree.right, huffman, path+"1")
	} else {
		huffman[tree.char] = path
	}
}

func walk(tree *node) map[string]string {
	huffman := make(map[string]string)
	walkRecursive(tree, huffman, "")
	return huffman
}

func pop(heap []*node) ([]*node, *node) {
	head := heap[0]
	heap[0], heap[len(heap)-1] = heap[len(heap)-1], heap[0]
	heap = heap[:len(heap)-1]
	downHeap(heap, 0)
	return heap, head
}

func downHeap(heap []*node, index int) {
	left, right := 2*index+1, 2*index+2
	switch {
	case left >= len(heap):
	case right >= len(heap):
		if heap[left].count < heap[index].count {
			heap[index], heap[left] = heap[left], heap[index]
			downHeap(heap, left)
		}
	default:
		var smaller int
		if heap[left].count < heap[right].count {
			smaller = left
		} else {
			smaller = right
		}
		if heap[smaller].count < heap[index].count {
			heap[index], heap[smaller] = heap[smaller], heap[index]
			downHeap(heap, smaller)
		}
	}
}

type node struct {
	left, right *node
	char        string
	count       int
}
