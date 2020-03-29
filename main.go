package main

import (
	pqueue "github.com/andela-sjames/priorityQueue"
)

func main() {
}

func dijstra(g []string, n int, s int) []string {
	// g - adjacency list of a weighted graph
	// n - the number of nodes in the graph
	// s - the index of the starting node ( 0 <= s < n )
	visited := make([]bool, n)
	distance := make([]int, n)

	// Set Min option to true for minheap
	minheap := pqueue.NewHeap(pqueue.Options{
		Min: true,
	})

	return []string{"B"}
}

// this should return the path for points distance covered.
func findShortestPath() []string {
	// g - adjacency list of a weighted graph
	// n - the number of nodes in the graph
	// s - the index of the starting node ( 0 <= s < n )

	return []string{"A"}
}
