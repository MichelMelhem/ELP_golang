package models

import "container/heap"

// Item represents a node in the priority queue
type Item struct {
	Node     int
	Priority int
	Index    int
}

// PriorityQueue implements a min-heap for Dijkstra
type PriorityQueue []Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(Item)
	item.Index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = Item{} // Avoid memory leak
	item.Index = -1   // For safety
	*pq = old[0 : n-1]
	return item
}

// Update modifies the priority of an item in the queue
func (pq *PriorityQueue) Update(item *Item, priority int) {
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
