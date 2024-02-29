package graph

func Degree(g IGraph, v int) int {
	degree := 0
	for range g.Adj(v) {
		degree++
	}

	return degree
}

func MaxDegree(g IGraph) int {
	ans := 0
	for v := 0; v < g.V(); v++ {
		if d := Degree(g, v); d > ans {
			ans = d
		}
	}

	return ans
}

func AvgDegree(g IGraph) float64 {
	return 2 * float64(g.E()) / float64(g.V())
}

func NumberOfSelfLoops(g IGraph) int {
	count := 0
	for v := 0; v < g.V(); v++ {
		for w := range g.Adj(v) {
			if v == w {
				count++
			}
		}
	}

	return count / 2
}
