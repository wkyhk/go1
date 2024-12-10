package leetcode

/*
743. 网络延迟时间
有 n 个网络节点，标记为 1 到 n。

给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点， wi 是一个信号从源节点传递到目标节点的时间。

现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。
链接：https://leetcode.cn/problems/network-delay-time/description/
*/
func networkDelayTime(times [][]int, n int, k int) int {
	graph := make([][]int, n+1)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, n+1)
		for j := 0; j < len(graph[i]); j++ {
			graph[i][j] = -1
		}
	}
	for _, time := range times {
		graph[time[0]][time[1]] = time[2]
	}
	return 0 //dijkstra(graph, k, n)
}
