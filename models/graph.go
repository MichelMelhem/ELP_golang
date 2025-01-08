package models

type Edge struct {
	to, weight int
}

type Graph struct {
	adjacencyList map[int][]Edge
}

func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[int][]Edge),
	}
}

func (g *Graph) AddEdge(from, to, weight int) {
	g.adjacencyList[from] = append(g.adjacencyList[from], Edge{to, weight})
}
