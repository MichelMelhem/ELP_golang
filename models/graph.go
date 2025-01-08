package models

// Création du type arrete avec un poids

type Edge struct {
	To, Weight int
}

// Création du type graphe avec une liste d'adjacence

type Graph struct {
	AdjacencyList map[int][]Edge
}

// Créé un nouveau graphe

func NewGraph() *Graph {
	return &Graph{
		AdjacencyList: make(map[int][]Edge),
	}
}

//ajouter une arrete au graphe

func (g *Graph) AddEdge(from, to, weight int) {
	g.AdjacencyList[from] = append(g.AdjacencyList[from], Edge{to, weight})
}

// Ajouter un noeud au graphe

func (g *Graph) AddNode(node int) {
	if _, exists := g.AdjacencyList[node]; !exists {
		g.AdjacencyList[node] = []Edge{}
	}
}
