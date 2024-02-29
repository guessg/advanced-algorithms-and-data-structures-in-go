package graph

import "github.com/guessg/aads-go/stack"

type DirectedCycle struct {
	marked  []bool
	edgeTo  []int  // edgeTo[v] = previous vertex on path to v
	cycle   []int  // 有向环中的所有顶点（如果存在）
	onStack []bool // 递归调用的栈上的所有顶点
}

func NewDirectedCycle(g IDirectedGraph) *DirectedCycle {
	dc := &DirectedCycle{
		marked:  make([]bool, g.V()),
		edgeTo:  make([]int, g.V()),
		onStack: make([]bool, g.V()),
	}

	for v := 0; v < g.V(); v++ {
		if !dc.marked[v] {
			dc.dfs(g, v)
		}
	}

	return dc
}

func (d *DirectedCycle) dfs(g IDirectedGraph, v int) {
	d.onStack[v] = true

	// marked 用于标记是否已经访问过，不会重新设置为 false
	d.marked[v] = true
	for w := range g.Adj(v) {
		if d.HasCycle() {
			return
		}

		if !d.marked[w] {
			d.edgeTo[w] = v
			d.dfs(g, w)
		} else if d.onStack[w] {
			// 有向环中的所有顶点
			for x := v; x != w; x = d.edgeTo[x] {
				d.cycle = append(d.cycle, x)
			}
			d.cycle = append(d.cycle, w)
			d.cycle = append(d.cycle, v)
		}
	}
	d.onStack[v] = false
}

func (d *DirectedCycle) HasCycle() bool {
	return len(d.cycle) != 0
}

func (d *DirectedCycle) Cycle() []int {
	return d.cycle
}

type DepthFirstOrder struct {
	marked      []bool
	pre         []int             // 递归之前顶点入队 queue
	post        []int             // 递归之后顶点入队 queue
	reversePost *stack.Stack[int] // 递归之后顶点入栈 stack
}

func NewDepthFirstOrder(g IDirectedGraph) *DepthFirstOrder {
	dfo := &DepthFirstOrder{
		marked:      make([]bool, g.V()),
		pre:         make([]int, 0),
		post:        make([]int, 0),
		reversePost: stack.NewStack[int](),
	}

	for v := 0; v < g.V(); v++ {
		if !dfo.marked[v] {
			dfo.dfs(g, v)
		}
	}
	return dfo
}

func (d *DepthFirstOrder) Pre() []int {
	return d.pre
}

func (d *DepthFirstOrder) Post() []int {
	return d.post
}

func (d *DepthFirstOrder) ReversePost() *stack.Stack[int] {
	return d.reversePost
}

func (d *DepthFirstOrder) dfs(g IDirectedGraph, v int) {
	d.pre = append(d.pre, v)

	d.marked[v] = true
	for w := range g.Adj(v) {
		if !d.marked[w] {
			d.dfs(g, w)
		}
	}

	d.post = append(d.post, v)
	d.reversePost.Push(v)
}

type Topological struct {
	order []int
}

func NewTopological(g IDirectedGraph) *Topological {
	dc := NewDirectedCycle(g)
	if !dc.HasCycle() {
		dfs := NewDepthFirstOrder(g)
		var order = make([]int, 0)
		// 处理逆后续顶点
		for !dfs.ReversePost().IsEmpty() {
			order = append(order, dfs.ReversePost().Pop())
		}
		return &Topological{order: order}
	}
	return nil
}

func (t *Topological) IsDAG() bool {
	return t.order != nil
}

func (t *Topological) Order() []int {
	return t.order
}
