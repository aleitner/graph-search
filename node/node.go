package node

import "fmt"

// Node in the graph
type Node struct {
	data byte
}

// String prints the node data
func (n *Node) String() string {
	return fmt.Sprintf("%v", n.data)
}

func NewNode(data byte) *Node {
	return &Node {
		data: data,
	}
}