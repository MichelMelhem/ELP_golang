# ELP_golang

## Description du programme

Le programme réalise un N-Dijkstra sur un graphe pondéré orienté de manière multithreadée. Il est accessible "as a service" via une requête TCP.

## Structure du programme

Le programme se compose de deux parties distinctes :
- **Le client** : S'exécute à partir du fichier `client.go`, situé dans le dossier `client`. Il permet d'envoyer le graphe vers le serveur et d'afficher le résultat.
- **Le serveur et l'algorithme** : Composés des fichiers restants, ils se lancent à partir du fichier `main.go`.

## Architecture du programme

Le programme est multithreadé de telle manière que, lors de l'exécution de N-Dijkstra, l'algorithme de Dijkstra est exécuté dans une go-routine différente pour chaque nœud du graphe.

## Setup

Ce projet a été programmé à partir de l'éditeur VS Code. Il contient une configuration qui permet de le démarrer en "un clic" à partir de l'interface graphique.
