package bfs_dfs

import (
	"fmt"
	"testing"
)

func TestUndirectedGraph_BreadthFirstSearch(t *testing.T) {
	g := fillGraph()
	g.BreadthFirstSearch(func(n *Node) {
		fmt.Printf("%v\n", n)
	})
}