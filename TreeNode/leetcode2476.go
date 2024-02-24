package TreeNode

import "sort"

/*
2476. 二叉搜索树最近节点查询
给你一个 二叉搜索树 的根节点 root ，和一个由正整数组成、长度为 n 的数组 queries 。

请你找出一个长度为 n 的 二维 答案数组 answer ，其中 answer[i] = [mini, maxi] ：

mini 是树中小于等于 queries[i] 的 最大值 。如果不存在这样的值，则使用 -1 代替。
maxi 是树中大于等于 queries[i] 的 最小值 。如果不存在这样的值，则使用 -1 代替。
返回数组 answer 。
链接：https://leetcode.cn/problems/closest-nodes-queries-in-a-binary-search-tree/description/
*/
func closestNodes(root *TreeNode, queries []int) [][]int {
	arr := []int{}
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		arr = append(arr, root.Val)
		dfs(root.Right)
	}

	dfs(root)
	res := make([][]int, len(queries))
	for i, val := range queries {
		maxVal, minVal := -1, -1
		index := sort.SearchInts(arr, val)
		if index < len(arr) {
			maxVal = arr[index]
			if arr[index] == val {
				minVal = arr[index]
				res[i] = []int{minVal, maxVal}
				continue
			}
		}
		if index != 0 {
			minVal = arr[index-1]
		}
		res[i] = []int{minVal, maxVal}
	}
	return res
}
