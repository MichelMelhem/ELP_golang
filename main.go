package main

import (
	"container/heap"
	"elp/models"
	"fmt"
	"math"
	"net"
	"runtime"
	"sync"
)

var MaxRoutines int = 10

func Dijkstra(g *models.Graph, source int) map[int]map[int]int {

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

		// Si la distance actuelle est plus grande que la distance déjà enregistrée, on passe
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

	return map[int]map[int]int{source: distances}
}

func worker(id int, jobs <-chan models.WorkerInput, results chan<- map[int]map[int]int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Le worker %d démarre\n", id)
	for j := range jobs {
		fmt.Printf("Worker %d commence le traitement du noeud source %d\n", id, j.Noeud)
		result := Dijkstra(j.Graph, j.Noeud)
		results <- result
		fmt.Printf("Worker %d a terminé le traitement du noeud source %d\n", id, j.Noeud)
	}
}

func NDisktra(graph *models.Graph) []map[string]interface{} {
	fmt.Println("Démarrage de NDisktra pour tout le graphe...")

	jobs := make(chan models.WorkerInput, MaxRoutines)
	results := make(chan map[int]map[int]int)
	var wg sync.WaitGroup

	fmt.Println("Initialisation des workers...")
	for w := 1; w <= MaxRoutines; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	go func() {
		fmt.Println("Envoi des jobs aux workers...")
		for source := range graph.AdjacencyList {
			jobs <- models.WorkerInput{Graph: graph, Noeud: source}
		}
		close(jobs)
	}()

	go func() {
		fmt.Println("Attente de la fin de tous les jobs...")
		wg.Wait() // Attente que tous les jobs soient terminés
		close(results)
	}()

	var resultJSON []map[string]interface{}
	fmt.Println("Récupération des résultats...")
	for result := range results {
		for source, distances := range result {
			// Ajoute de la réponse au résultat
			fmt.Printf("Ajout des résultats pour le noeud source %d\n", source)
			resultJSON = append(resultJSON, map[string]interface{}{
				"source":    source,
				"distances": distances,
			})
		}
	}

	fmt.Println("Tous les résultats ont été récupérés et ajoutés.")
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

	MaxRoutines = runtime.NumCPU()

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
