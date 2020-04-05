package main

import (
	"fmt"
	"math"
	"strconv"

	pqueue "github.com/andela-sjames/priorityQueue"
)

type adjList []*linkedlist

func main() {

	// create adj list
	nodezero := newlist("nodezero")
	nodeone := newlist("nodeone")
	nodetwo := newlist("nodetwo")
	nodethree := newlist("nodethree")
	nodefour := newlist("nodefour")

	adjlist := []*linkedlist{nodezero, nodeone, nodetwo, nodethree, nodefour}

	n := 4
	s := 0

	res := dijstra(adjlist, n, s)
	fmt.Println(res)
}

type node struct {
	data   int
	weight int
	next   *node
}

func newlist(name string) *linkedlist {
	return &linkedlist{
		name: name,
	}
}

type linkedlist struct {
	name string
	head *node
}

func (l *linkedlist) addNode(data, weight int) error {
	n := &node{
		data:   data,
		weight: weight,
	}
	if l.head == nil {
		l.head = n
	} else {
		currentNode := l.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = n
	}
	return nil
}

func dijstra(g adjList, n int, s int) []string {
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
