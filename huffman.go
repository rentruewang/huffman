package main

import (
	"container/heap"
)

// Any can be anything!
type Any = interface{}

// HuffmanNode is actually a normal node, but huffman.
// It is created to aid the process of creating Huffman trees.
type HuffmanNode struct {
	// left points to the left HuffmanNode.
	left *HuffmanNode
	// right points to the right HuffmanNode.
	right *HuffmanNode
	// token holds what the node represents.
	// If the node is an internal node, the the token is rune(0), representing '\0'.
	// Or else the token is representative of a particular token in the document.
	// Please see Wikipedia for the definition of a token
	token rune
	// Number of times the token is counted in the document
	count int
}

// hasLeft shows if the node has a left child.
func (hn HuffmanNode) hasLeft() bool { return hn.left != nil }

// hasRight shows if the node has a right child.
func (hn HuffmanNode) hasRight() bool { return hn.right != nil }

// ValidToken shows if a token is valid for non-package level access.
// An invalid token is an implementation detail and should not be visible to users.
func (hn HuffmanNode) ValidToken() bool { return hn.token != rune(0) }

// Token represented by the HuffmanNode. The field is not modified ever after the creation of the huffman node.
func (hn HuffmanNode) Token() rune { return hn.token }

// Count is the number of times a rune is present in the document.
// If the token is a ValidToken, the count is not modified after the creation of the huffman node.
func (hn HuffmanNode) Count() int { return hn.count }

// HuffmanTree is, guess what, a pointer to the root HuffmanNode of the tree!
// HuffmanTree is an alias because if it weren't, the spec (in the following line)
// https://golang.org/ref/spec#Method_declarations
// says that it you can't define methods on it.
// The reason for the outrageous decision is actually pretty logical,
// see the link below
// https://groups.google.com/g/golang-nuts/c/qf76N-uDcHA/m/DTCDNgaF_p4J
type HuffmanTree = *HuffmanNode

// MakeHuffmanTree creates a completely new HuffmanTree from a string
func MakeHuffmanTree(content string) HuffmanTree {
	// wordCount will serve as a multiset to count every occurrence of a rune
	wordCount := make(map[rune]int)

	for _, char := range content {
		// If key doesn't exist, the map defaults to returning 0
		// And since this map is used to count numbers of encounters
		// Simply put ++ works fine as the value 'inserted' into the map is 1
		wordCount[char]++
	}

	// The reason I'm not setting the capacity to length of wordCount
	// is that if I do so, the next line would have to maintain a separate integer
	// just to count the index of huffmanList.
	// Using a preset-length of 0 makes code more readable by using append only
	list := MakeHuffmanList(0)
	for word, count := range wordCount {
		// The left and right nodes are currently nil because they are not yet connected.
		// Will do so in the next line.
		list.Append(HuffmanNode{left: nil, right: nil, token: word, count: count})
	}

	// The list will be built into a min heap
	heap.Init(&list)

	// Run for list.Len() - 1 iterations
	for i := list.Len() - 1; i > 0; i-- {
		// Retrieve the two nodes where the counts are the smallest
		// Also, I'm pretty sure they are still HuffmanNode's
		first := heap.Pop(&list).(HuffmanNode)
		second := heap.Pop(&list).(HuffmanNode)

		// Why create a parent with `Count` the sum of its children?
		// Well, this way the count of the root of a tree
		// will be the sum of all its leaf nodes,
		// since all internal nodes being the sum of their children,
		// will then propagate the sums to their parents.
		// And the above condition holds for all sub-trees in the tree.
		merged := HuffmanNode{
			left:  &first,
			right: &second,
			token: rune(0),
			count: first.count + second.count,
		}

		// Then insert the new node back to the original list
		// The list looks like a list of trees after the first heap.Push
		heap.Push(&list, merged)
	}

	// After n-1 merges (remember total=len(list)-1 in the for loop?)
	// There would only be one element left in the list
	if list.Len() != 1 {
		panic("unreachable")
	}

	// HuffmanTree is a *HuffmanNode
	// The 'type conversion' is just for readability
	return HuffmanTree(&list[0])
}

// Huffman generates huffman codes by recursively calling GenerateByPath
func (hf HuffmanTree) Huffman() map[string]string {
	dict := make(map[string]string)

	// The root node begins with an empty dictionary and no path at all
	// As the tree is traversed by the following function,
	// paths are assigned to the node and that is the node's Huffman code!
	hf.GenerateByPath(dict, "")

	return dict
}

// GenerateByPath generates the huffman codes for an existing HuffmanTree
func (ht HuffmanTree) GenerateByPath(dict map[string]string, path string) {
	// If hf is an internal node
	// then there is nothing to store because characters will only be present in the leaf nodes
	if ht.ValidToken() {
		// type string is easier to display and use
		token := string(ht.token)
		dict[token] = path
	}

	// In a huffman tree the physical layout of the tree is related to how the nodes are named.
	// A node is named by following the path from the root to it.
	// Every left visit (visit left child) a '0' is appended to the code
	// Every right visit (visit right child) a '1' is appended to the code

	if ht.hasLeft() {
		ht.left.GenerateByPath(dict, path+"0")
	}

	if ht.hasRight() {
		ht.right.GenerateByPath(dict, path+"1")
	}
}

// Huffman is a convenient method to create
// a huffman tree and the compute the huffman codes
// Basically, it takes in a document string,
// then compute the huffman codes for you
func Huffman(content string) map[string]string {
	ht := MakeHuffmanTree(content)
	return ht.Huffman()
}

// HuffmanList is used for building up the HuffmanTree
// So it's very short lived. RIP
type HuffmanList []HuffmanNode

// MakeHuffmanList creates a new HuffmanList with given length.
func MakeHuffmanList(length int) HuffmanList { return make(HuffmanList, length) }

// Append adds a new HuffmanNode to HuffmanList
// The reason copy is used here is that the node is not that big anyways,
// and it makes the API simpler
func (hl *HuffmanList) Append(node HuffmanNode) { *hl = append(*hl, node) }

// Index returns the value of HuffmanList at a certain index
// I created this method because if someone were to extend this package,
// HuffmanList without this method would be useless
func (hl HuffmanList) Index(idx int) HuffmanNode { return hl[idx] }

// The following functions Len, Less, Swap are defined on HuffmanList,
// but if an interface requires these three methods,
// both HuffmanList and *HuffmanList implement the interface.
// It's Go's magic.
// https://golangbyexample.com/pointer-vs-value-receiver-method-golang/
// So what we effectively that sort package can use both HuffmanList and *HuffmanList,
// and heap package can only use *HuffmanList.

// Len is defined for sort, heap package.
func (hl HuffmanList) Len() int { return len(hl) }

// Less is defined for sort, heap package.
// It is desired that the heap is a min heap. So Less(i, j) uses the normal`<` operator.
// It would be clearer why we want a min heap in creating the HuffmanTree.
func (hl HuffmanList) Less(i, j int) bool { return hl[i].count < hl[j].count }

// Swap is defined for sort, heap package.
func (hl HuffmanList) Swap(i, j int) { hl[i], hl[j] = hl[j], hl[i] }

// Push is defined for heap package
func (hl *HuffmanList) Push(elem Any) {
	if node, ok := elem.(HuffmanNode); ok {
		*hl = append(*hl, node)
		return
	}
	// Unreachable because I trust the standard library not to do stupid things
	panic("unreachable")
}

// Pop is defined for heap package
func (hl *HuffmanList) Pop() (last Any) {
	list := *hl

	// The last index
	index := len(list) - 1

	// Takes out the last element and shorten existing list
	*hl, last = list[:index], list[index]

	return last
}
