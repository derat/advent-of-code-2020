package main

import (
	"fmt"

	"github.com/derat/advent-of-code/lib"
)

func main() {
	input := lib.InputInts("2018/8")
	root, cnt := readNode(input)
	lib.AssertEq(cnt, len(input))
	fmt.Println(root.metasum())
	fmt.Println(root.value())
}

type node struct {
	children []*node
	metadata []int
}

// readNode reads a node (and any subnodes) from the beginning of data.
// It returns the root node and the number of values that were consumed.
func readNode(data []int) (*node, int) {
	read := 2
	nchild, nmeta := data[0], data[1]
	n := &node{make([]*node, nchild), make([]int, nmeta)}

	for i := 0; i < nchild; i++ {
		child, cnt := readNode(data[read:])
		n.children[i] = child
		read += cnt
	}
	for i := 0; i < nmeta; i++ {
		n.metadata[i] = data[read]
		read++
	}

	return n, read
}

func (n *node) metasum() int {
	sum := lib.Sum(n.metadata...)
	for _, c := range n.children {
		sum += c.metasum()
	}
	return sum
}

func (n *node) value() int {
	if len(n.children) == 0 {
		return lib.Sum(n.metadata...)
	}
	var val int
	for _, idx := range n.metadata {
		idx-- // convert from 1-indexed to 0-indexed
		if idx >= 0 && idx < len(n.children) {
			val += n.children[idx].value()
		}
	}
	return val
}
