package main

import (
	"container/heap"
	"elp/models"
	"fmt"
	"math"
)

// Dijkstra calculates shortest paths from a source node
func Dijkstra(g *models.Graph, source int) map[int]int {
	distances := make(map[int]int)
	for node := range g.AdjacencyList {
		distances[node] = math.MaxInt32
	}
	distances[source] = 0

	pq := &models.PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, models.Item{Node: source, Priority: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(models.Item)
		currentNode := current.Node
		currentDist := current.Priority

		if currentDist > distances[currentNode] {
			continue
		}

		for _, edge := range g.AdjacencyList[currentNode] {
			newDist := distances[currentNode] + edge.Weight
			if newDist < distances[edge.To] {
				distances[edge.To] = newDist
				heap.Push(pq, models.Item{Node: edge.To, Priority: newDist})
			}
		}
	}

	return distances
}

func runDijkstraExample() {
	graph := models.NewGraph()
	graph.AddEdge(1, 2, 4)
	graph.AddEdge(1, 3, 2)
	graph.AddEdge(2, 3, 5)
	graph.AddEdge(2, 4, 10)
	graph.AddEdge(3, 4, 3)

	source := 1
	distances := Dijkstra(graph, source)

	fmt.Printf("Shortest distances from node %d:\n", source)
	for node, distance := range distances {
		fmt.Printf("Node %d: %d\n", node, distance)
	}
}

func main() {
	runDijkstraExample()
}
