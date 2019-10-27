package stack

import (
	"bfs-dfs/node"
	"sync"
)

// Stack of Nodes used for Depth first Search
type Stack struct {
	nodes []node.Node
	mtx sync.RWMutex
}

// NewStack returns an empty Stack
func New() *Stack {
	return &Stack{
		nodes: []node.Node{},
	}
}

// Add a node to the stack
func (s *Stack) Add(n node.Node) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.nodes = append(s.nodes, n)
}

// Take a node off the stack
func (s *Stack) Take() *node.Node {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if len(s.nodes) == 0 {
		return nil
	}

	pulledNode := s.nodes[len(s.nodes) - 1]
	s.nodes = s.nodes[0:len(s.nodes) - 1]

	return &pulledNode
}

// Front checks the node that is at the front of the queue
func (s *Stack) Front() *node.Node {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	if len(s.nodes) == 0 {
		return nil
	}

	return &s.nodes[len(s.nodes) - 1]
}


// Size of the stack
func (s *Stack) Size() int {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	return len(s.nodes)
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	return len(s.nodes) == 0
}