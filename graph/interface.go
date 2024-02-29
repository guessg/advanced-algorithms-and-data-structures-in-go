package graph

type IGraph interface {
	V() int
	E() int
	AddEdge(v, w int) // v - w || v -> w
	Adj(v int) []int
	String() string
}

type IDirectedGraph interface {
	IGraph
	Reverse() IDirectedGraph
}

type IPath interface {
	// HasPathTo Is there a path between s and v
	HasPathTo(v int) bool
	// PathTo dfs between s and v
	PathTo(v int) []int
}

type ISearch interface {
	// Marked Is there a path between s and v
	Marked(v int) bool
	// Count Number of vertices connected to s
	Count() int
}

type IConnectedComponents interface {
	// Connected Are v and w connected
	Connected(v, w int) bool
	// Id Component identifier for v
	Id(v int) int
	// Count Number of connected components
	Count() int
}
