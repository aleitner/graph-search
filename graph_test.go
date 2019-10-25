package bfs_dfs

import (
	"testing"
)

func fillGraph() *UndirectedGraph {
	g := NewUndirectedGraph()

	nA := Node{1}
	nB := Node{2}
	nC := Node{3}
	nD := Node{4}
	nE := Node{5}
	nF := Node{6}
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