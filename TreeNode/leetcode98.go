package TreeNode

import "math"

/*
98. 验证二叉搜索树
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
有效 二叉搜索树定义如下：
节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
链接：https://leetcode.cn/problems/validate-binary-search-tree
*/
func IsValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}
	var Dfs func(*TreeNode, int, int) bool
	Dfs = func(node *TreeNode, max int, min int) bool {
		if node == nil {
			return true
		}
		if node.Val >= max || node.Val <= min {
			return false
		}
		return Dfs(node.Right, max, node.Val) && Dfs(node.Left, node.Val, min)
	}
	return Dfs(root, math.MaxInt64, math.MinInt64)
}
