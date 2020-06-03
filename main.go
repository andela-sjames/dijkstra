package main

import (
	"fmt"
	"strconv"

	pqueue "github.com/andela-sjames/priorityQueue"
)

type adjList []*linkedlist

func main() {

	nodezero := newlist("nodezero")
	nodezero.addNode(1, 4)
	nodezero.addNode(2, 1)

	nodeone := newlist("nodeone")
	nodeone.addNode(3, 1)

	nodetwo := newlist("nodetwo")
	nodetwo.addNode(1, 2)
	nodetwo.addNode(3, 5)

	nodethree := newlist("nodethree")
	nodethree.addNode(4, 3)

	nodefour := newlist("nodefour")
	adjlist := []*linkedlist{nodezero, nodeone, nodetwo, nodethree, nodefour}

	n := len(adjlist)
	s := 0
	e := 4

	dist, path := findShortestPath(adjlist, n, s, e)
	fmt.Println(dist, path)
}

type node struct {
	vertex int
	weight int
	next   *node
}

func newlist(name string) *linkedlist {
	return &linkedlist{
		name: name,
	}
}

func reverseSlice(a []int) []int {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

type linkedlist struct {
	name string
	head *node
}

func (l *linkedlist) addNode(vertex, weight int) {
	n := &node{
		vertex: vertex,
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
}

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
		distance[i] = 1000 // some high value ? revisit
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

func findShortestPath(g adjList, n int, s int, e int) ([]int, []int) {
	// g - adjacency list of a weighted graph
	// n - the number of nodes in the graph
	// s - the index of the starting node ( 0 <= s < n )
	// e - the index of the end node (0 <= e < n )
	dist, prev := dijstra(g, n, s, e)
	path := make([]int, 0)

	if dist[e] == 1000 { // need this to be infinity later
		return path, nil
	}

	// start from the end and loop all the way back to the beginning.
	for pointer := e; pointer != 0; pointer = prev[pointer] {
		path = append(path, pointer)
	}
	sp := reverseSlice(path)

	return dist, sp
}
