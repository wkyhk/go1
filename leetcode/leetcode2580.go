package leetcode

import "sort"

/*
2580. 统计将重叠区间合并成组的方案数
给你一个二维整数数组 ranges ，其中 ranges[i] = [starti, endi] 表示 starti 到 endi 之间（包括二者）的所有整数都包含在第 i 个区间中。

你需要将 ranges 分成 两个 组（可以为空），满足：

每个区间只属于一个组。
两个有 交集 的区间必须在 同一个 组内。
如果两个区间有至少 一个 公共整数，那么这两个区间是 有交集 的。

比方说，区间 [1, 3] 和 [2, 5] 有交集，因为 2 和 3 在两个区间中都被包含。
请你返回将 ranges 划分成两个组的 总方案数 。由于答案可能很大，将它对 109 + 7 取余 后返回。
链接：https://leetcode.cn/problems/count-ways-to-group-overlapping-ranges/description/
*/
func countWays(ranges [][]int) int {
	const mod = int(1e9 + 7)
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	n := len(ranges)
	res := int64(1)
	for i := 0; i < n; {
		r := ranges[i][1]
		j := i + 1
		for j < n && ranges[j][0] <= r {
			r = max(r, ranges[j][1])
			j++
		}
		res = (res * 2) % int64(mod)
		i = j
	}
	return int(res)
}
