package leetcode

/*
2368.受限条件下可到达节点的数目
现有一棵由 n 个节点组成的无向树，节点编号从 0 到 n - 1 ，共有 n - 1 条边。

给你一个二维整数数组 edges ，长度为 n - 1 ，其中 edges[i] = [ai, bi] 表示树中节点 ai 和 bi 之间存在一条边。另给你一个整数数组 restricted 表示 受限 节点。

在不访问受限节点的前提下，返回你可以从节点 0 到达的 最多 节点数目。

注意，节点 0 不 会标记为受限节点。
链接：https://leetcode.cn/problems/reachable-nodes-with-restrictions/description/
*/
func reachableNodes(n int, edges [][]int, restricted []int) int {
	isRestricted := make([]int, n)
	for _, x := range restricted {
		isRestricted[x] = 1
	}
	g := make([][]int, n)
	for _, v := range edges {
		g[v[0]] = append(g[v[0]], v[1])
		g[v[1]] = append(g[v[1]], v[0])
	}

	cnt := 0
	var dfs func(int, int)
	dfs = func(x, f int) {
		cnt++
		for _, y := range g[x] {
			if y != f && isRestricted[y] != 1 {
				dfs(y, x)
			}
		}
	}
	dfs(0, -1)
	return cnt
}
