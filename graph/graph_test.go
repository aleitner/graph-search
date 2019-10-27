package graph

import (
	"bfs-dfs/node"
	"bfs-dfs/queue"
	"bfs-dfs/stack"
	"fmt"
	"testing"
)

func fillGraph() *UndirectedGraph {
	g := NewUndirectedGraph()

	nA := node.Node{1}
	nB := node.Node{2}
	nC := node.Node{3}
	nD := node.Node{4}
	nE := node.Node{5}
	nF := node.Node{6}
	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nC)
	g.AddEdge(&nB, &nE)
	g.AddEdge(&nC, &nE)
	g.AddEdge(&nE, &nF)
	g.AddEdge(&nD, &nA)

	return g
}

func TestAdd(t *testing.T) {
	g := fillGraph()
	g.Print()
}

func TestBreadthFirstSearch(t *testing.T) {
	g := fillGraph()
	g.Search(queue.New(), func(n *node.Node) {
		fmt.Printf("%v\n", n)
	})
}

func TestDepthFirstSearch(t *testing.T) {
	g := fillGraph()
	g.Search(stack.New(), func(n *node.Node) {
		fmt.Printf("%v\n", n)
	})
}