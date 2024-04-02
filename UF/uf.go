package UF

type UF struct {
	parent []int
	count  int
}

func NewUF(n int) *UF {
	uf := &UF{
		count: n,
	}

	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}

	return uf
}

func (u *UF) Union(p, q int) {
	rootP := u.Find(p)
	rootQ := u.Find(q)
	if rootP == rootQ {
		return
	}

	u.parent[rootQ] = rootP
	u.count--
}

func (u *UF) Find(x int) int {
	// 路径压缩
	if u.parent[x] != x {
		u.parent[x] = u.Find(u.parent[x])
	}

	return u.parent[x]
}

func (u *UF) Connected(p, q int) bool {
	rootP := u.Find(p)
	rootQ := u.Find(q)
	return rootP == rootQ
}

func (u *UF) Count() int {
	return u.count
}
