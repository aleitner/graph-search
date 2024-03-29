package graph

import (
	"bfs-dfs/node"
	"fmt"
	"sync"
)

// Graph of nodes and edges
type UndirectedGraph struct {
	nodes []*node.Node
	edges map[node.Node][]*node.Node
	mtx sync.RWMutex
}

type SearchDataStructure interface {
	Add(t node.Node)
	Take() *node.Node
	Front() *node.Node
	IsEmpty() bool
	Size() int
}

// NewUndirectedGraph creates a new graph with no direction
func NewUndirectedGraph() *UndirectedGraph {
	return &UndirectedGraph{
		edges: make(map[node.Node][]*node.Node),
	}
}

// AddNode to the Graph
func (g* UndirectedGraph) AddNode(n *node.Node) {
	g.mtx.Lock()
	defer g.mtx.Unlock()

	// TODO: Add check for if Node already exists

	g.nodes = append(g.nodes, n)
}

// AddEdge between two nodes
func (g *UndirectedGraph) AddEdge(n1, n2 *node.Node) {
	g.mtx.Lock()
	defer g.mtx.Unlock()

	// TODO: Add check for if Node already exists as an edge

	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
}

// Print the graph
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

func (g *UndirectedGraph) Search(dataStructure SearchDataStructure, callback func(*node.Node)) {
	g.mtx.RLock()
	defer g.mtx.RUnlock()

	if len(g.nodes) == 0 {
		return
	}

	entranceNode := g.nodes[0]
	dataStructure.Add(*entranceNode)
	visited := make(map[*node.Node]bool)

	for {
		if dataStructure.IsEmpty() {
			return
		}

		node := dataStructure.Take()
		visited[node] = true
		near := g.edges[*node]

		for i := 0; i < len(near); i++ {
			nearbyNode := near[i]
			if !visited[nearbyNode] {
				dataStructure.Add(*nearbyNode)
				visited[nearbyNode] = true
			}
		}

		if callback != nil {
			callback(node)
		}
	}
}