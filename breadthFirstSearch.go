package bfs_dfs

func (g *UndirectedGraph) BreadthFirstSearch(callback func(*Node)) {
	g.mtx.RLock()
	defer g.mtx.RUnlock()

	if len(g.nodes) == 0 {
		return
	}

	q := NewQueue()
	entranceNode := g.nodes[0]
	q.Enqueue(*entranceNode)
	visited := make(map[*Node]bool)

	for {
		if q.IsEmpty() {
			return
		}

		node := q.Dequeue()
		visited[node] = true
		near := g.edges[*node]

		for i := 0; i < len(near); i++ {
			nearbyNode := near[i]
			if !visited[nearbyNode] {
				q.Enqueue(*nearbyNode)
				visited[nearbyNode] = true
			}
		}

		if callback != nil {
			callback(node)
		}
	}

}