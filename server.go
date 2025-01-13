package main

import (
	"bufio"
	"elp/models"
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

//server TCP qui reçois la connexion

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	// Read input until newline
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(conn, "Error reading input:", err)
		return
	}

	// Trim whitespace
	input = strings.TrimSpace(input)

	// Parse the JSON into a graph
	graph, err := models.ParseJSONToGraph(input)
	if err != nil {
		fmt.Fprintln(conn, "Error parsing JSON:", err)
		return
	}

	// Compute Dijkstra results
	response, err := json.Marshal(NDisktra(graph))
	if err != nil {
		fmt.Fprintln(conn, "Error generating response JSON:", err)
		return
	}

	// Send the response back to the client
	conn.Write(append(response, '\n')) // Ensure newline for client
}
