package graph

type Graph struct {
	vertices map[int][]int
}

func New() *Graph {
	return &Graph{
		vertices: make(map[int][]int),
	}
}

func (g *Graph) AddEdge(v, w int) {
	g.vertices[v] = append(g.vertices[v], w)
	g.vertices[w] = append(g.vertices[w], v)
}

func (g *Graph) BFS(start int) []int {
	visited := make(map[int]bool)
	queue := []int{start}
	result := []int{}

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		if !visited[v] {
			visited[v] = true
			result = append(result, v)

			for _, neighbor := range g.vertices[v] {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
				}
			}
		}
	}

	return result
}

func (g *Graph) DFS(start int) []int {
	visited := make(map[int]bool)
	result := []int{}

	var dfs func(v int)
	dfs = func(v int) {
		visited[v] = true
		result = append(result, v)

		for _, neighbor := range g.vertices[v] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	dfs(start)
	return result
}

func (g *Graph) HasCycle() bool {
	visited := make(map[int]bool)

	var hasCycleUtil func(v int, parent int) bool
	hasCycleUtil = func(v int, parent int) bool {
		visited[v] = true

		for _, neighbor := range g.vertices[v] {
			if !visited[neighbor] {
				if hasCycleUtil(neighbor, v) {
					return true
				}
			} else if neighbor != parent {
				return true
			}
		}

		return false
	}

	for v := range g.vertices {
		if !visited[v] {
			if hasCycleUtil(v, -1) {
				return true
			}
		}
	}

	return false
}

func (g *Graph) GetNeighbors(v int) []int {
	return g.vertices[v]
}
