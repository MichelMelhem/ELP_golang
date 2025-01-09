package main

import (
	"bufio"
	"elp/models"
	"encoding/json"
	"fmt"
	"net"
)

//server TCP qui reçois la connexion

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(conn, "Error reading input:", err)
		return
	}

	graph, err := models.ParseJSONToGraph(input)
	if err != nil {
		fmt.Fprintln(conn, "Error parsing JSON:", err)
		return
	}

	response, err := json.Marshal(NDisktra(graph))
	if err != nil {
		fmt.Fprintln(conn, "Error generating response JSON:", err)
		return
	}

	conn.Write(response)
}
