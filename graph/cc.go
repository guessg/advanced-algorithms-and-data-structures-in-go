package graph

// CC connected components
type CC struct {
	marked []bool
	id     []int
	count  int
}

func NewCC(g IGraph) *CC {
	cc := &CC{
		marked: make([]bool, g.V()),
		id:     make([]int, g.V()),
		count:  0,
	}
	for s := 0; s < g.V(); s++ {
		if !cc.marked[s] {
			cc.dfs(g, s)
			cc.count++
		}
	}
	return cc
}

func (c *CC) dfs(g IGraph, s int) {
	c.marked[s] = true
	c.id[s] = c.count
	for w := range g.Adj(s) {
		if !c.marked[w] {
			c.dfs(g, w)
		}
	}
}

func (c *CC) Connected(v, w int) bool {
	return c.id[v] == c.id[w]
}

func (c *CC) Id(v int) int {
	return c.id[v]
}

func (c *CC) Count() int {
	return c.count
}
