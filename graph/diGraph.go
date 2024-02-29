package graph

type Digraph struct {
	v   int
	e   int
	adj [][]int
}

func NewDigraph(v int) *Digraph {
	g := &Digraph{
		v:   v,
		adj: make([][]int, v),
	}

	for i := 0; i < v; i++ {
		g.adj[i] = make([]int, 0)
	}

	return g
}

// V returns the number of vertices in the graph
func (g *Digraph) V() int {
	return g.v
}

// E returns the number of edges in the graph
func (g *Digraph) E() int {
	return g.e
}

func (g *Digraph) AddEdge(v, w int) {
	g.e++
	g.adj[v] = append(g.adj[v], w)
}

// Adj returns the vertices adjacent to v
func (g *Digraph) Adj(v int) []int {
	return g.adj[v]
}

// String returns the string representation of the graph
func (g *Digraph) String() string {
	return ""
}

func (g *Digraph) Reverse() *Digraph {
	r := NewDigraph(g.V())
	for v := 0; v < g.V(); v++ {
		for _, w := range g.Adj(v) {
			r.AddEdge(w, v)
		}
	}
	return r
}
