package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
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

	queue := make(nodelist, 0)
	for key, val := range wordCount {
		queue = append(queue, node{char: key, count: val})
	}

	heap.Init(&queue)
	for i, total := 0, len(queue)-1; i < total; i++ {
		first := heap.Pop(&queue).(node)
		second := heap.Pop(&queue).(node)
		n := node{left: &first, right: &second, char: "", count: first.count + second.count}
		heap.Push(&queue, n)
	}

	huffman := walk(&queue[0])
	fmt.Println(huffman)

	total := 0
	for key, str := range huffman {
		total += len(str) * wordCount[key]
	}
	fmt.Println("cost", total)
}

func walkRecursive(tree *node, huffman map[string]string, path string, prl *parallel) {
	defer prl.Done()
	if tree.char == "" {
		prl.Add(2)
		go walkRecursive(tree.left, huffman, path+"0", prl)
		go walkRecursive(tree.right, huffman, path+"1", prl)
	} else {
		prl.Lock()
		huffman[tree.char] = path
		prl.Unlock()
	}
}

func walk(tree *node) map[string]string {
	huffman := make(map[string]string)
	var prl parallel
	prl.Add(1)
	go walkRecursive(tree, huffman, "", &prl)
	prl.Wait()
	return huffman
}

type parallel struct {
	sync.WaitGroup
	sync.Mutex
}

type node struct {
	left, right *node
	char        string
	count       int
}

type nodelist []node

// Len for heap package
func (n nodelist) Len() int { return len(n) }

// Less for heap package
func (n nodelist) Less(i, j int) bool { return n[i].count < n[j].count }

// Swap for heap package
func (n nodelist) Swap(i, j int) { n[i], n[j] = n[j], n[i] }

// Push for heap package
func (n *nodelist) Push(elem interface{}) {
	*n = append(*n, elem.(node))
}

// Pop for heap package
func (n *nodelist) Pop() interface{} {
	old := *n
	ln := len(old)
	x := old[ln-1]
	*n = old[:ln-1]
	return x
}
