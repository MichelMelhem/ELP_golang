package main

import (
	"container/heap"
	"elp/models"
	"fmt"
	"math"
	"net"
	"sync"
)

func Dijkstra(g *models.Graph, source int, results chan<- map[int]map[int]int, wg *sync.WaitGroup) {
	defer wg.Done() // n'envoie pas le signal que le thread a terminé tant que la fonction n'est pas fini

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

	// envoie le résultat du thread
	results <- map[int]map[int]int{source: distances}
}

// appelle Disktra sur tout les neud en créant à chaque fois un nouveau thread
func NDisktra(graph *models.Graph) []map[string]interface{} {
	results := make(chan map[int]map[int]int)
	var wg sync.WaitGroup

	// Launch Dijkstra for each node in the graph
	for source := range graph.AdjacencyList {
		wg.Add(1) //on ajoute 1 au waitgroup dès qu'on lance un nouveau thread
		go Dijkstra(graph, source, results, &wg)
	}

	// Close the results channel when all routines are done
	go func() {
		wg.Wait()
		close(results)
	}()

	var resultJSON []map[string]interface{}
	// Récupère les resultats
	for result := range results {
		for source, distances := range result {
			// Ajoute de la réponse au résultat
			resultJSON = append(resultJSON, map[string]interface{}{
				"source":    source,
				"distances": distances,
			})
		}
	}
	fmt.Print(resultJSON)
	return resultJSON
}

func runDijkstraExample() {
	graph := models.NewGraph()

	graph.AddNode(1)
	graph.AddNode(2)
	graph.AddNode(3)
	graph.AddNode(4)

	graph.AddEdge(1, 2, 4)
	graph.AddEdge(1, 3, 2)
	graph.AddEdge(2, 3, 5)
	graph.AddEdge(3, 4, 3)
	graph.AddEdge(2, 4, 10)

	NDisktra(graph)

}

func main() {
	// run manuellement l'algorithme
	//runDijkstraExample()
	//démare le serveur tcp pour utiliser Dkistra as a service
	ln, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Server is running on port 12345...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
