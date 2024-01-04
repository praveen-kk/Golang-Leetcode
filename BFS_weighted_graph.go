package main

import "fmt"

type Edge struct {
	To     int
	Weight int
}
type Graph struct {
	Vertices     int
	AdjacentList map[int][]Edge
}

func newGraph(vertices int) *Graph {
	return &Graph{
		Vertices:     vertices,
		AdjacentList: make(map[int][]Edge),
	}
}
func (g *Graph) AddEdge(src, dest, weight int) {
	g.AdjacentList[src] = append(g.AdjacentList[src], Edge{
		To:     dest,
		Weight: weight,
	})
	g.AdjacentList[dest] = append(g.AdjacentList[dest], Edge{
		To:     src,
		Weight: weight,
	},
	)
}

//shortestPathFind finds the shortest path from a source vertex to every other vertex using BFS

func (g *Graph) ShortestPathFind(startVertex int) map[int]int {

	distances := make(map[int]int)
	fmt.Println("\nline 37  AdjacentList: ", g.AdjacentList)
	for v := range g.AdjacentList {
		distances[v] = -1
	}
	distances[startVertex] = 0
	fmt.Println("line 42 distances : ", distances)
	queue := []int{startVertex}
	fmt.Println("line 43 queue: ", queue)

	for len(queue) > 0 {
		currentVertex := queue[0]
		fmt.Println("line 46 currentVertex: ", currentVertex)
		queue = queue[1:]
		fmt.Println("line 49 queue: ", queue)

		for _, edge := range g.AdjacentList[currentVertex] {
			fmt.Println("line 51 g.AdjacentList[currentVertex]:  ", g.AdjacentList[currentVertex])
			fmt.Printf("line 54 distances[currentVertex]: %d\n  edge.Weight : %d ", distances[currentVertex], edge.Weight)
			newDistance := distances[currentVertex] + edge.Weight
			fmt.Println("line 53 newDistance:  ", newDistance)

			if distances[edge.To] == -1 || newDistance < distances[edge.To] {
				distances[edge.To] = newDistance
				fmt.Println("line 60 distances[edge.To]:  ", distances[edge.To])
				fmt.Println("line 61 distances : ", distances)
				queue = append(queue, edge.To)

				fmt.Println("line 64 queue: ", queue)

			}
		}
	}
	fmt.Println(distances)
	return distances
}

func main() {

	g := newGraph(9)

	g.AddEdge(0, 1, 3)
	g.AddEdge(0, 2, 1)
	g.AddEdge(0, 3, 2)
	g.AddEdge(3, 4, 2)
	g.AddEdge(4, 5, 2)
	g.AddEdge(4, 8, 8)
	g.AddEdge(2, 5, 11)
	g.AddEdge(0, 9, 22)

	fmt.Print(g)
	startVertex := 0
	g.ShortestPathFind(startVertex)

	distances := g.ShortestPathFind(startVertex)
	for v, distance := range distances {
		fmt.Printf("shortest distance from vertex %d - %d: %d\n", startVertex, v, distance)
	}
}
