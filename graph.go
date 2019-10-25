package bfs_dfs

import (
	"fmt"
	"sync"
)

type UndirectedGraph struct {
	nodes []*Node
	edges map[Node][]*Node
	mtx sync.RWMutex
}

type Node struct {
	data byte
}

func NewUndirectedGraph() *UndirectedGraph {
	return &UndirectedGraph{
		edges: make(map[Node][]*Node),
	}
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.data)
}

func (g* UndirectedGraph) AddNode(n *Node) {
	g.mtx.Lock()
	defer g.mtx.Unlock()

	// TODO: Add check for if Node already exists

	g.nodes = append(g.nodes, n)
}

func (g *UndirectedGraph) AddEdge(n1, n2 *Node) {
	g.mtx.Lock()
	defer g.mtx.Unlock()

	// TODO: Add check for if Node already exists as an edge

	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
}

func (g *UndirectedGraph) Print() {
	g.mtx.RLock()
	defer g.mtx.RUnlock()

	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + " -> "
		near := g.edges[*g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}

	fmt.Println(s)
}
