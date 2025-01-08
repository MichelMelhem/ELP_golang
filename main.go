package main

import (
	"container/heap"
	"fmt"
	"math"
)

// PriorityQueue implements a min-heap for Dijkstra
type PriorityQueue []Item

// Item represents a node in the priority queue
type Item struct {
	node     int
	priority int
	index    int
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = Item{} // Avoid memory leak
	item.index = -1   // For safety
	*pq = old[0 : n-1]
	return item
}

// Update modifies the priority of an item in the queue
func (pq *PriorityQueue) Update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

// Dijkstra calculates shortest paths from a source node
func Dijkstra(g *Graph, source int) map[int]int {
	distances := make(map[int]int)
	for node := range g.adjacencyList {
		distances[node] = math.MaxInt32
	}
	distances[source] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Item{node: source, priority: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(Item)
		currentNode := current.node
		currentDist := current.priority

		if currentDist > distances[currentNode] {
			continue
		}

		for _, edge := range g.adjacencyList[currentNode] {
			newDist := distances[currentNode] + edge.weight
			if newDist < distances[edge.to] {
				distances[edge.to] = newDist
				heap.Push(pq, Item{node: edge.to, priority: newDist})
			}
		}
	}

	return distances
}

func main() {
	graph := NewGraph()
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
