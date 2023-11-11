package leetcode

/*
765. 情侣牵手
n对情侣坐在连续排列的2n个座位上，想要牵到对方的手。

计算最少交换座位的次数，以便每对情侣可以并肩坐在一起。
人和座位由一个整数数组 row 表示，其中 row[i] 是坐在第 i 个座位上的人的 ID。情侣们按顺序编号，第一对是 (0, 1)，第二对是 (2, 3)，以此类推，最后一对是 (2n-2, 2n-1)。

返回 最少交换座位的次数，以便每对情侣可以并肩坐在一起。 每次交换可选择任意两人，让他们站起来交换座位。

链接：https://leetcode.cn/problems/couples-holding-hands/description/
*/

func minSwapsCouples(row []int) int {
	n := len(row)
	N := n / 2
	uf := ConstructorUnionFind765(N)
	for i := 0; i < n; i += 2 {
		uf.union(row[i]/2, row[(i+1)]/2)
	}
	return N - uf.count
}

type UnionFind765 struct {
	parent []int
	count  int
}

func ConstructorUnionFind765(total int) *UnionFind765 {
	p := make([]int, total)
	for i := 0; i < total; i++ {
		p[i] = i
	}
	return &UnionFind765{
		parent: p,
		count:  total,
	}
}
func (u *UnionFind765) union(x, y int) {
	xp := u.find(x)
	yp := u.find(y)
	if xp == yp {
		return
	}
	u.parent[xp] = yp
	u.count--
}
func (u *UnionFind765) find(x int) int {
	root := x
	for root != u.parent[root] {
		root = u.parent[root]
	}
	i, j := x, 0
	for root != u.parent[i] {
		j = u.parent[i]
		u.parent[i] = root
		i = j
	}
	return root
}
