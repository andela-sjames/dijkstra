package main

import (
	"fmt"

	"github.com/andela-sjames/dijkstra/utils"
)

func main() {

	nodezero := utils.NewList("nodezero")
	nodezero.AddNode(1, 4)
	nodezero.AddNode(2, 1)

	nodeone := utils.NewList("nodeone")
	nodeone.AddNode(3, 1)

	nodetwo := utils.NewList("nodetwo")
	nodetwo.AddNode(1, 2)
	nodetwo.AddNode(3, 5)

	nodethree := utils.NewList("nodethree")
	nodethree.AddNode(4, 3)

	nodefour := utils.NewList("nodefour")
	adjlist := []*utils.LinkedList{nodezero, nodeone, nodetwo, nodethree, nodefour}

	n := len(adjlist)
	s := 0
	e := 4

	dist, path := utils.FindShortestPath(adjlist, n, s, e)
	fmt.Println("Distance graph: ", dist)
	fmt.Println("Shortest Path: ", path)
	fmt.Println("Shortest Distance: ", dist[4])
}
