package main

import (
	"fmt"
)

type Graph struct {
	Vertices      int
	adjacencyList map[int][]int
}

func newGraph(vertices int) *Graph {

	return &Graph{
		Vertices:      vertices,
		adjacencyList: make(map[int][]int),
	}

}

func (g *Graph) bfs(startVertex int) {
	visited := make(map[int]bool)
	queue := []int{startVertex}

	visited[startVertex] = true
	//queue.PushBack(startVertex)

	for len(queue) > 0 {
		currentVertex := queue[0]
		queue = queue[1:]
		fmt.Print(currentVertex, " ")

		for _, neighbour := range g.adjacencyList[currentVertex] {
			if !visited[neighbour] {
				visited[neighbour] = true
				queue = append(queue, neighbour)
			}
		}
	}

}

func (g *Graph) addEdges(src, dest int) {
	g.adjacencyList[src] = append(g.adjacencyList[src], dest)
	g.adjacencyList[dest] = append(g.adjacencyList[dest], src)
}

func main() {

	v := newGraph(9)
	v.addEdges(0, 1)
	v.addEdges(0, 2)
	v.addEdges(0, 5)
	v.addEdges(1, 3)
	v.addEdges(1, 7)
	v.addEdges(0, 7)
	v.addEdges(1, 8)
	v.addEdges(5, 9)
	v.addEdges(9, 10)

	fmt.Println("bfs from vertex <vertex number>: ")
	v.bfs(9)

}
