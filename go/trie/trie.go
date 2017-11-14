package trie

import (
	_ "fmt"
)

type Node struct {
	child map[int]*Node

	value    interface{}
	hasValue bool
}

func NewNode() *Node {
	return &Node{
		child: make(map[int]*Node),
	}
}

func (n *Node) Load(key string, val interface{}) {
	current := n.create(key)
	current.hasValue = true
	current.value = val
}

func (n *Node) create(key string) *Node {
	current := n
	for _, c := range key {
		child, ok := current.child[int(c)]
		if !ok {
			child = NewNode()
			current.child[int(c)] = child
		}
		current = child
	}

	return current
}

func (n *Node) find(key string) *Node {
	current := n

	for _, c := range key {
		if child, ok := current.child[int(c)]; !ok {
			break
		} else {
			current = child
		}
	}
	return current
}

func (n *Node) Find(key string) (value interface{}, find bool) {
	node := n.find(key)
	if node == nil {
		return nil, false
	}
	if node.hasValue == false {
		return nil, false
	}

	return node.value, node.hasValue
}
