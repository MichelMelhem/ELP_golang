package models

type Edge struct {
	To, Weight int
}

type Graph struct {
	AdjacencyList map[int][]Edge
}

func NewGraph() *Graph {
	return &Graph{
		AdjacencyList: make(map[int][]Edge),
	}
}

func (g *Graph) AddEdge(from, to, weight int) {
	g.AdjacencyList[from] = append(g.AdjacencyList[from], Edge{to, weight})
}
