package bfs_dfs

import "sync"

// Queue of Nodes used for Breadth first Search
type Queue struct {
	nodes []Node
	mtx sync.RWMutex
}

// Create a New Queue
func NewQueue() *Queue {
	return &Queue{
		nodes: []Node{},
	}
}

// Enqueue a node
func (q *Queue) Enqueue(t Node) {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	q.nodes = append(q.nodes, t)
}

// Dequeue the node that is next in queue
func (q *Queue) Dequeue() *Node {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	if len(q.nodes) == 0 {
		return nil
	}

	node := q.nodes[0]
	q.nodes = q.nodes[1:len(q.nodes)]

	return &node
}

// Front checks the node that is at the front of the queue
func (q *Queue) Front() *Node {
	q.mtx.RLock()
	defer q.mtx.RUnlock()

	if len(q.nodes) == 0 {
		return nil
	}

	return &q.nodes[0]
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	q.mtx.RLock()
	defer q.mtx.RUnlock()

	return len(q.nodes) == 0
}

// Size of the queue
func (q *Queue) Size() int {
	q.mtx.RLock()
	defer q.mtx.RUnlock()

	return len(q.nodes)
}