package TreeNode

import "sort"

/*
2583. 二叉树中的第 K 大层和
给你一棵二叉树的根节点 root 和一个正整数 k 。

树中的 层和 是指 同一层 上节点值的总和。

返回树中第 k 大的层和（不一定不同）。如果树少于 k 层，则返回 -1 。

注意，如果两个节点与根节点的距离相同，则认为它们在同一层。
链接:https://leetcode.cn/problems/kth-largest-sum-in-a-binary-tree/description/
*/
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	ans := make([]int64, 0)
	var bfs func(root *TreeNode, level int)
	bfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if len(ans) < level {
			ans = append(ans, int64(root.Val))
		} else {
			ans[level-1] += int64(root.Val)
		}
		bfs(root.Left, level+1)
		bfs(root.Right, level+1)
	}
	bfs(root, 1)
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] > ans[j]
	})
	if len(ans) < k {
		return -1
	}
	return ans[k-1]
}
