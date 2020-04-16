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

	n := 5
	s := 0

	dist, prev := dijstra(adjlist, n, s)
	fmt.Println(dist, prev)
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

func dijstra(g adjList, n int, s int) ([]int, []int) {
	// g - adjacency list of a weighted graph
	// n - the number of nodes in the graph
	// s - the index of the starting node ( 0 <= s < n )
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

		stringIndex, min := minheap.Poll()
		integerIndex, _ := strconv.Atoi(stringIndex)

		// current node is integerIndex
		visited[integerIndex] = true

		// optimization to ignore stale index stale
		// (index, min_dis) pair
		if distance[integerIndex] < min {
			continue
		}

		// loop through all the neighbours of
		// the current node
		cn := g[integerIndex].head
		for cn != nil {

			if visited[cn.vertex] {
				continue
			}

			newdist := distance[integerIndex] + cn.weight
			if newdist < distance[cn.vertex] {
				previous[cn.vertex] = integerIndex
				distance[cn.vertex] = newdist
				minheap.InsertPriority(strconv.Itoa(cn.vertex), newdist)
			}

			if cn.next == nil {
				break
			}
			cn = cn.next
		}
	}
	return distance, previous
}

func findShortestPath() {

}
