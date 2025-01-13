package main

import (
	"bufio"
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

func main() {
	// Connexion au serveurs
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Création du graph
	graphJSON := `{
		"0": [{"To": 1, "Weight": 2}, {"To": 2, "Weight": 4}],
		"1": [{"To": 2, "Weight": 1}, {"To": 3, "Weight": 7}],
		"2": [{"To": 3, "Weight": 3}],
		"3": []
	}`

	graphJSON = strings.ReplaceAll(graphJSON, "\n", "")
	graphJSON = strings.ReplaceAll(graphJSON, "\t", "")

	// Envoie du graph sous format json au serveur
	fmt.Fprintln(conn, graphJSON)

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
