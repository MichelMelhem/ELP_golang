package models

import (
	"encoding/json"
	"fmt"
)

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

// converti un string en entier
func atoi(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}

/*
Exemple de json valide pour la fonction ParseJSONToGraph:
    {
		"0": [{"To": 1, "Weight": 2}, {"To": 2, "Weight": 4}],
		"1": [{"To": 2, "Weight": 1}, {"To": 3, "Weight": 7}],
		"2": [{"To": 3, "Weight": 3}],
		"3": []
	}
*/

func ParseJSONToGraph(jsonData string) (*Graph, error) {
	var rawGraph map[string][]map[string]int
	if err := json.Unmarshal([]byte(jsonData), &rawGraph); err != nil {
		return nil, err
	}

	graph := NewGraph()
	for node, edges := range rawGraph {
		nodeID := atoi(node)
		for _, edge := range edges {

			graph.AddEdge(nodeID, edge["To"], edge["Weight"])
		}
	}
	return graph, nil
}
