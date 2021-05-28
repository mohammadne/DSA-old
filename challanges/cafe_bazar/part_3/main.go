package main

import "fmt"

type Node struct {
	character byte
	children  []Node
	endCounts int
}

func (n *Node) satisfyPrefix() bool {
	if len(n.children) == 0 && n.endCounts <= 1 {
		return false
	}

	if n.endCounts > 0 && len(n.children) > 0 || n.endCounts > 1 {
		return true
	} else {
		satisfy := false

		for _, child := range n.children {
			if child.satisfyPrefix() {
				satisfy = true
				break
			}
		}

		return satisfy
	}
}

func (n *Node) addChar(char byte) *Node {
	node := Node{character: char}
	n.children = append(n.children, node)
	length := len(n.children)

	return &(n.children[length-1])
}

func (n *Node) contains(char byte) *Node {
	var contains *Node = nil

	for index := 0; index < len(n.children); index++ {
		if n.children[index].character == char {
			contains = &(n.children[index])
			break
		}
	}

	return contains
}

type Trie struct {
	root *Node
}

func (t *Trie) hasPrefixString() bool {
	return t.root.satisfyPrefix()
}

func (trie *Trie) insert(str string) {
	length := len(str)
	current := trie.root

	for index := 0; index < length; index++ {
		if contains := current.contains(str[index]); contains != nil {
			current = contains
		} else {
			current = current.addChar(str[index])
		}

		if index == length-1 {
			current.endCounts += 1
		}
	}
}

func main() {
	input := []string{
		"abcd",
		"abcd",
	}

	trie := Trie{
		root: &Node{},
	}

	for _, str := range input {
		trie.insert(str)
	}

	fmt.Printf("has prefix: %t", trie.hasPrefixString())

}
