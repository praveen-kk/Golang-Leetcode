package main

/*
In below depth first search traversal is implemented using Recursive method
*/
import "fmt"

type Graph struct {
	vertices      int
	adjacencyList map[int][]int
}

func newGraph(Vertices int) *Graph {
	return &Graph{
		vertices:      Vertices,
		adjacencyList: make(map[int][]int),
	}

}

func (g *Graph) addEdge(src, dest int) {

	g.adjacencyList[src] = append(g.adjacencyList[src], dest)
	g.adjacencyList[dest] = append(g.adjacencyList[dest], src)

}

func (g *Graph) DFS(startVertex int, visted map[int]bool) {
	if visted == nil {
		visted = make(map[int]bool)
	}

	visted[startVertex] = true
	fmt.Printf("%d ", startVertex)

	for _, neighbour := range g.adjacencyList[startVertex] {

		//fmt.Printf("line 38 neighbour %d : , g.adjacencyList[startVertex] %d: - startVertex %d \n", neighbour, g.adjacencyList[startVertex], startVertex)
		if !visted[neighbour] {
			//fmt.Printf("line 40 visted[neighbour] %d\n : ", visted[neighbour])
			g.DFS(neighbour, visted)
		}
	}
}

func main() {

	g := newGraph(8)
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(0, 3)
	g.addEdge(2, 5)
	g.addEdge(4, 6)
	g.addEdge(1, 4)
	g.addEdge(4, 8)

	fmt.Println(g)
	startVertex := 0
	visted := make(map[int]bool)
	g.DFS(startVertex, visted)
	fmt.Println()
}
