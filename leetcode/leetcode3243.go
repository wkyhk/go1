package leetcode
/*
3243. 新增道路查询后的最短距离 I 
给你一个整数 n 和一个二维整数数组 queries。

有 n 个城市，编号从 0 到 n - 1。初始时，每个城市 i 都有一条单向道路通往城市 i + 1（ 0 <= i < n - 1）。

queries[i] = [ui, vi] 表示新建一条从城市 ui 到城市 vi 的单向道路。每次查询后，你需要找到从城市 0 到城市 n - 1 的最短路径的长度。

返回一个数组 answer，对于范围 [0, queries.length - 1] 中的每个 i，answer[i] 是处理完前 i + 1 个查询后，从城市 0 到城市 n - 1 的最短路径的长度。
链接：https://leetcode.cn/problems/shortest-distance-after-road-addition-queries-i/description/
 */
 func shortestDistanceAfterQueries(n int, queries [][]int) []int {
    neighbors := make([][]int, n)
    for i := 0; i < n - 1; i++ {
        neighbors[i] = append(neighbors[i], i + 1)
    }
    var res []int
    for _, query := range queries {
        neighbors[query[0]] = append(neighbors[query[0]], query[1])
        res = append(res, bfs(n, neighbors))
    }
    return res
}

func bfs(n int, neighbors [][]int) int {
    dist := make([]int, n)
    for i := 1; i < n; i++ {
        dist[i] = -1
    }
    q := []int{0}
    for len(q) > 0 {
        x := q[0]
        q = q[1:]
        for _, y := range neighbors[x] {
            if dist[y] >= 0 {
                continue
            }
            q = append(q, y)
            dist[y] = dist[x] + 1
        }
    }
    return dist[n - 1]
}
