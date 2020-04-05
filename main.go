package main

import (
	"fmt"
	"math"
	"strconv"

	pqueue "github.com/andela-sjames/priorityQueue"
)

func main() {

	// create a proper adj list
	//
	adjList := [][]int{
		{1, 6, 8},
		{0, 4, 6, 9},
		{4, 6},
		{4, 5, 8},
	}

	n := 10
	s := 0

	res := dijstra(adjList, n, s)
	fmt.Println(res)
}

type node struct {
	data   int
	weight int
	next   *node
}

func dijstra(g [][]int, n int, s int) []string {
	// g - adjacency list of a weighted graph
	// n - the number of nodes in the graph
	// s - the index of the starting node ( 0 <= s < n )
	visited := make([]bool, n)
	distance := make([]int, n)

	for i := range visited {
		visited[i] = false
	}

	for i := range distance {
		distance[i] = int(math.Inf(+1))
	}

	distance[s] = 0
	// Set Min option to true for minheap
	minheap := pqueue.NewHeap(pqueue.Options{
		Min: true,
	})

	minheap.InsertPriority(string(s), 0)

	for {
		if minheap.Length() == 0 {
			break
		}

		// while it's not 0
		index, priority := minheap.Poll()
		i1, _ := strconv.Atoi(index)
		visited[i1] = true

		// loop through all the neighbours of
		// the current node

	}

	return []string{"B"}
}

// this should return the path for points distance covered.
func findShortestPath() []string {

	return []string{"A"}
}
