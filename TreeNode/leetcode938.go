package TreeNode

/*
938. 二叉搜索树的范围和
给定二叉搜索树的根结点 root，返回值位于范围 [low, high] 之间的所有结点的值的和。
https://leetcode.cn/problems/range-sum-of-bst/
*/
func rangeSumBST(root *TreeNode, low int, high int) int {
	ans := 0
	var bft func(*TreeNode)
	bft = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Val >= low && node.Val <= high {
			ans += node.Val

		}

		bft(node.Right)
		bft(node.Left)
	}
	bft(root)
	return ans
}
