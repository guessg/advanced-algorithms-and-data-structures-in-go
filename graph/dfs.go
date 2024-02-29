package graph

// DFS 在访问其中一个顶点时：
// 将它标记为已访问；
// 递归地访问它的所有没有被标记过的邻居顶点。
type DFS struct {
	marked []bool // v 和 s 连通吗？
	count  int    // 与 s 连通的顶点总数
}

func NewDFSSearch(g IGraph, s int) *DFS {
	search := &DFS{
		marked: make([]bool, g.V()),
		count:  0,
	}
	search.dfs(g, s)
	return search
}

func (d *DFS) dfs(g IGraph, s int) {
	d.marked[s] = true
	d.count++
	for w := range g.Adj(s) {
		if !d.marked[w] {
			d.dfs(g, w)
		}
	}
}

func (d *DFS) Marked(v int) bool {
	return d.marked[v]
}

func (d *DFS) Count() int {
	return d.count
}

type DirectedDFS struct {
	marked []bool
	count  int
}

func NewDirectedDFSSearch(g IDirectedGraph, s int) *DirectedDFS {
	search := &DirectedDFS{
		marked: make([]bool, g.V()),
		count:  0,
	}
	search.dfs(g, s)
	return search
}

func (d *DirectedDFS) dfs(g IDirectedGraph, s int) {
	d.marked[s] = true
	d.count++
	for w := range g.Adj(s) {
		if !d.marked[w] {
			d.dfs(g, w)
		}
	}
}

func (d *DirectedDFS) Marked(v int) bool {
	return d.marked[v]
}
