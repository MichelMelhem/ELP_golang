package main

import (
	"bufio"
	"elp/models"
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

// crée une structure pour mapper la réponse du serveur
type DijkstraResult struct {
	Source    int         `json:"source"`
	Distances map[int]int `json:"distances"`
}

func generateLargeGraph(numNodes, maxEdgesPerNode int) *models.Graph {
	graph := models.NewGraph()

	for i := 0; i < numNodes; i++ {
		graph.AddNode(i)

		for j := 1; j <= maxEdgesPerNode && i+j < numNodes; j++ {
			graph.AddEdge(i, i+j, i+j)
		}
	}

	return graph
}

func main() {
	// Connexion au serveurs
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Création du graph
	graphJSON, _ := json.Marshal(generateLargeGraph(25, 4))

	graphAsAString := string(graphJSON)

	graphAsAString = strings.ReplaceAll(graphAsAString, "\n", "")
	graphAsAString = strings.ReplaceAll(graphAsAString, "\t", "")

	// Envoie du graph sous format json au serveur
	fmt.Fprintln(conn, graphAsAString)

	// Lis la réponse du serveur
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading server response:", err)
		return
	}

	// Extrait la réponse du serveur qui est au format json
	var results []DijkstraResult
	err = json.Unmarshal([]byte(response), &results)
	if err != nil {
		fmt.Println("Error parsing server response:", err)
		fmt.Println("Raw response:", response)
		return
	}

	// Imprime le résultat
	fmt.Println("Dijkstra Results:")
	for _, result := range results {
		fmt.Printf("Depuis le noeud %d: %v\n", result.Source, result.Distances)
	}
}
