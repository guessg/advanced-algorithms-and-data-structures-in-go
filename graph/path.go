package graph

type DepthFirstPath struct {
	marked []bool
	edgeTo []int // edgeTo[v] = previous vertex on path to v
	s      int   // source vertex
}

func NewDepthFirstPath(g IGraph, s int) *DepthFirstPath {
	path := &DepthFirstPath{
		marked: make([]bool, g.V()),
		edgeTo: make([]int, g.V()),
		s:      s,
	}
	path.dfs(g, s)
	return path
}

func (d *DepthFirstPath) dfs(g IGraph, s int) {
	d.marked[s] = true
	for _, w := range g.Adj(s) {
		if !d.marked[w] {
			d.edgeTo[w] = s
			d.dfs(g, w)
		}
	}
}

func (d *DepthFirstPath) HasPathTo(v int) bool {
	return d.marked[v]
}

func (d *DepthFirstPath) PathTo(v int) []int {
	if !d.HasPathTo(v) {
		return nil
	}

	path := make([]int, 0)
	for x := v; x != d.s; x = d.edgeTo[x] {
		path = append(path, x)
	}
	path = append(path, d.s)

	return path
}

type BreadthFirstPath struct {
	marked []bool
	edgeTo []int // edgeTo[v] = previous vertex on path to v
	s      int   // source vertex
}

func NewBreadthFirstPath(g IGraph, s int) *BreadthFirstPath {
	bfs := &BreadthFirstPath{
		marked: make([]bool, g.V()),
		edgeTo: make([]int, g.V()),
		s:      s,
	}
	bfs.dfs(g, s)
	return bfs
}

func (b *BreadthFirstPath) dfs(g IGraph, s int) {
	queue := make([]int, 0)
	b.marked[s] = true
	queue = append(queue, s)
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		for _, w := range g.Adj(v) {
			if !b.marked[w] {
				b.edgeTo[w] = v
				b.marked[w] = true
				queue = append(queue, w)
			}
		}
	}
}

func (b *BreadthFirstPath) HasPathTo(v int) bool {
	return b.marked[v]
}

func (b *BreadthFirstPath) PathTo(v int) []int {
	if !b.HasPathTo(v) {
		return nil
	}

	path := make([]int, 0)
	for x := v; x != b.s; x = b.edgeTo[x] {
		path = append(path, x)
	}
	path = append(path, b.s)

	return path
}
