package main

import (
	"container/heap"
	"sort"
)

// Any is anything.
type Any = interface{}

// HuffmanNode a node in a Huffman tree.
type HuffmanNode struct {
	// left points to the left HuffmanNode.
	left *HuffmanNode
	// right points to the right HuffmanNode.
	right *HuffmanNode
	// token stored by a node.
	// If the node is internal, then the token is rune(0), representing '\0'.
	// Or else the token is a particular token in the document.
	token rune
	// count is the number of times the token is present in the document.
	count int
}

// hasLeft shows if the node has a left child.
func (hn HuffmanNode) hasLeft() bool { return hn.left != nil }

// hasRight shows if the node has a right child.
func (hn HuffmanNode) hasRight() bool { return hn.right != nil }

// ValidToken shows if a token is valid for non-package level access.
func (hn HuffmanNode) ValidToken() bool { return hn.token != rune(0) }

// Token represented by the HuffmanNode.
// The field is not modified ever after the creation of the huffman node.
func (hn HuffmanNode) Token() rune { return hn.token }

// Count is the number of times a rune is present in the document.
// If the token is a `ValidToken`, the count is not modified after the creation of the huffman node.
func (hn HuffmanNode) Count() int { return hn.count }

// HuffmanTree is a pointer to the root HuffmanNode of the tree.
// HuffmanTree cannot be a newtype because if it is, the spec (in the following line)
// https://golang.org/ref/spec#Method_declarations
// says that it you can't define methods on it.
// https://groups.google.com/g/golang-nuts/c/qf76N-uDcHA/m/DTCDNgaF_p4J
type HuffmanTree = *HuffmanNode

// MakeHuffmanTree creates a new HuffmanTree from a string.
func MakeHuffmanTree(content string) HuffmanTree {
	// wordCount is a multiset.
	wordCount := make(map[rune]int)

	for _, char := range content {
		// If key doesn't exist, the map defaults to returning 0.
		wordCount[char]++
	}

	list := MakeHuffmanList(0)
	for word, count := range wordCount {
		list.Append(HuffmanNode{left: nil, right: nil, token: word, count: count})
	}

	heap.Init(&list)

	for i := list.Len() - 1; i > 0; i-- {
		// Retrieve the smallest (in `count`) two nodes.
		first := heap.Pop(&list).(HuffmanNode)
		second := heap.Pop(&list).(HuffmanNode)

		// Every parent having `count` the sum of its children's `count`s.
		// will ensure that the node of a tree would be the sum of its leaves.
		merged := HuffmanNode{
			left:  &first,
			right: &second,
			token: rune(0),
			count: first.count + second.count,
		}

		heap.Push(&list, merged)
	}

	// After n-1 merges there would only be one element left in the list.
	if list.Len() != 1 {
		panic("unreachable")
	}

	return HuffmanTree(&list[0])
}

// Huffman generates huffman codes by recursively appending '0' or '1' to the huffman node's code.
func (ht HuffmanTree) Huffman() map[string]string {
	// The root node's path == "".
	dict := make(map[string]string)
	ht.genByPath(dict, "")
	return dict
}

// genByPath generates the huffman codes for an existing HuffmanTree.
func (ht HuffmanTree) genByPath(dict map[string]string, path string) {
	// Tokens will only be present in the leaf nodes.
	if ht.ValidToken() {
		token := string(ht.token)
		dict[token] = path
	}

	// A node is named by following the path from the root to it.
	// Every left visit (visit left child) a '0' is appended to the code.
	// Every right visit (visit right child) a '1' is appended to the code.

	if ht.hasLeft() {
		ht.left.genByPath(dict, path+"0")
	}

	if ht.hasRight() {
		ht.right.genByPath(dict, path+"1")
	}
}

// Huffman takes in a document string and compute the huffman codes into a `map[string]string`.
func Huffman(content string) map[string]string {
	ht := MakeHuffmanTree(content)
	return ht.Huffman()
}

// HuffmanList is used for building up the HuffmanTree.
type HuffmanList []HuffmanNode

// HuffmanList is compliant with `sort.Interface`.
var _ sort.Interface = HuffmanList{}

// *HuffmanList is compliant with `sort.Interface`.
var _ sort.Interface = (*HuffmanList)(nil)

// *HuffmanList is compliant with `heap.Interface`.
var _ heap.Interface = (*HuffmanList)(nil)

// MakeHuffmanList creates a new HuffmanList with given length.
func MakeHuffmanList(length int) HuffmanList { return make(HuffmanList, length) }

// Append adds a new HuffmanNode to HuffmanList.
func (hl *HuffmanList) Append(node HuffmanNode) { *hl = append(*hl, node) }

// Index returns the value of HuffmanList at a certain index.
func (hl HuffmanList) Index(idx int) HuffmanNode { return hl[idx] }

// The following functions Len, Less, Swap are defined on HuffmanList,
// but if an interface requires these three methods,
// both HuffmanList and *HuffmanList implement the interface.
// https://golangbyexample.com/pointer-vs-value-receiver-method-golang/
// End result is: sort package can use both HuffmanList and *HuffmanList,
// and heap package can only use *HuffmanList.

// Len shows the length of the list. It is defined for sort, heap package.
func (hl HuffmanList) Len() int { return len(hl) }

// Less compares two nodes based on `Count()`. It is defined for sort, heap package.
func (hl HuffmanList) Less(i, j int) bool { return hl[i].count < hl[j].count }

// Swap swaps two elements. It is defined for sort, heap package.
func (hl HuffmanList) Swap(i, j int) { hl[i], hl[j] = hl[j], hl[i] }

// Push adds an element to the tail of a list. It is defined for heap package.
func (hl *HuffmanList) Push(elem Any) {
	hl.Append(elem.(HuffmanNode))
}

// Pop pops the last element from a list. It is defined for heap package.
func (hl *HuffmanList) Pop() (elem Any) {
	list := *hl
	last := len(list) - 1

	// Takes out the last element and shorten existing list.
	*hl, elem = list[:last], list[last]
	return
}
