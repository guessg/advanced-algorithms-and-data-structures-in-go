package graph

type IGraph interface {
	V() int
	E() int
	AddEdge(v, w int)
	Adj(v int) []int
	String() string
}
