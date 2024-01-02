package main

import "fmt"

/*
depth first search using stack
*/

type Edge struct {
	To     int
	Weight int
}

type graph struct {
	Vertices int
	Adjlist  map[int][]Edge
}

func newGraph(vertices int) *graph {
	return &graph{
		Vertices: vertices,
		Adjlist:  make(map[int][]Edge),
	}

}

func (g *graph) AddEdges(src, dest, weight int) {
	g.Adjlist[src] = append(g.Adjlist[src], Edge{
		To:     dest,
		Weight: weight,
	})
	g.Adjlist[dest] = append(g.Adjlist[dest], Edge{
		To:     src,
		Weight: weight,
	})

}

func DFSwithStack(g *graph, startVertex int) {

	visited := make(map[int]bool)
	stack := []int{startVertex}
	fmt.Printf("line 43 %d stack ", stack)

	for len(stack) > 0 {
		currentVertex := stack[len(stack)-1] //current vertex moves to 1 after [5 2 1]
		fmt.Printf("line 47 stack : %d\n  ", stack)
		stack = stack[:len(stack)-1] //52

		if !visited[currentVertex] {
			visited[currentVertex] = true
			fmt.Printf("%d ", currentVertex) //current vertex is printed 1

			for i := len(g.Adjlist[currentVertex]) - 1; i >= 0; i-- {
				neighbour := g.Adjlist[currentVertex][i]
				if !visited[neighbour.To] {
					stack = append(stack, neighbour.To)
					fmt.Printf("line 58 stack : %d\n ", stack)
				}
			}

		}

	}

}

func main() {

	g := newGraph(6)
	g.AddEdges(0, 1, 2)
	g.AddEdges(0, 2, 3)
	g.AddEdges(0, 5, 5)
	g.AddEdges(2, 3, 2)
	g.AddEdges(3, 4, 2)

	fmt.Println(g)

	startVertex := 0
	DFSwithStack(g, startVertex)
	fmt.Println()
}
