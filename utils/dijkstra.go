package utils

import (
	"math"
	"strconv"

	pqueue "github.com/andela-sjames/priorityQueue"
)

/** 
dijkstra function in it's base form takes a directed acyclic
and uses a naive approach - a non Indexed Priority Queue (IPQ)
to determine the shortest distance from the start of a graph to
the end of the graph.

It returns a slice of the shortest distance and a slice of the previous
node traversed to get to the end of the graph.
**/
func dijstra(g adjList, n int, s int, e int) ([]int, []int) {
	// g - adjacency list of a weighted graph
	// n - the number of nodes in the graph
	// s - the index of the starting node ( 0 <= s < n )
	// e - the index of the end node ( 0 <= e < n )
	visited := make([]bool, n)
	distance := make([]int, n)

	// keep track of the previous node we took
	// to get to the current node
	previous := make([]int, n)

	for i := range visited {
		visited[i] = false
	}

	for i := range distance {
		distance[i] = math.MaxInt64
	}

	distance[s] = 0
	// Set Min option to true for minheap
	minheap := pqueue.NewHeap(pqueue.Options{
		Min: true,
	})

	minheap.InsertPriority(string(s), 0)

	for minheap.Length() != 0 {

		stringAtIndex, min := minheap.Poll()
		integerAtIndex, _ := strconv.Atoi(stringAtIndex)

		// current node is integerAtIndex
		visited[integerAtIndex] = true

		// optimization to ignore stale index
		// (index, min_dis) pair
		if distance[integerAtIndex] < min {
			continue
		}

		// loop through all the neighbours of
		// the current node
		cn := g[integerAtIndex].head
		for cn != nil {

			if visited[cn.vertex] {
				continue
			}

			newdist := distance[integerAtIndex] + cn.weight
			if newdist < distance[cn.vertex] {
				previous[cn.vertex] = integerAtIndex
				distance[cn.vertex] = newdist
				minheap.InsertPriority(strconv.Itoa(cn.vertex), newdist)
			}

			if cn.next == nil {
				break
			}
			cn = cn.next
		}

		// Optimise here to stop early.
		if integerAtIndex == e {
			return distance, previous
		}

	}
	return distance, previous
}
