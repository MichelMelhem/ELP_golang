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

	// Lis la données d'entrée
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(conn, "Erreur lors de la lecture de l'entrée:", err)
		return
	}

	// Enlève les espaces qui sont éventuellement dans la data d'entrée
	input = strings.TrimSpace(input)

	// Parse le json vers une structure de Graph
	graph, err := models.ParseJSONToGraph(input)
	if err != nil {
		fmt.Fprintln(conn, "Erreur :", err)
		return
	}

	// Compute Dijkstra results
	response, err := json.Marshal(NDisktra(graph))
	if err != nil {
		fmt.Fprintln(conn, "Erreur lors de l'exportation du résultat au format Json :", err)
		return
	}

	//Envoie la réponse
	conn.Write(append(response, '\n'))
}
