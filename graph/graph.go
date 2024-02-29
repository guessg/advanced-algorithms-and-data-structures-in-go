package graph

type Graph struct {
	v   int
	e   int
	adj [][]int
}

func NewGraph(v int, edges [][2]int) *Graph {
	g := &Graph{
		v:   v,
		adj: make([][]int, 0),
	}

	for _, edge := range edges {
		v, w := edge[0], edge[1]
		g.AddEdge(v, w)
	}

	return g
}

// V returns the number of vertices in the graph
func (g *Graph) V() int {
	return g.v
}

// E returns the number of edges in the graph
func (g *Graph) E() int {
	return g.e
}

// AddEdge adds an edge between v and w
func (g *Graph) AddEdge(v, w int) {
	g.e++
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
}

// Adj returns the vertices adjacent to v
func (g *Graph) Adj(v int) []int {
	return g.adj[v]
}

// String returns the string representation of the graph
func (g *Graph) String() string {
	return ""
}
