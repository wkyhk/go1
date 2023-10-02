package TreeNode

/*
236. 二叉树的最近公共祖先
https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var node *TreeNode
	var F func(*TreeNode, *TreeNode, *TreeNode) bool
	F = func(root *TreeNode, p *TreeNode, q *TreeNode) bool {
		if root == nil {
			return false
		}
		b := root.Val == p.Val || root.Val == q.Val
		c := F(root.Right, p, q)
		d := F(root.Left, p, q)
		if b {
			if c || d {
				node = root
			}
		} else {
			if c && d {
				node = root
			}
		}
		return b || c || d
	}
	F(root, p, q)
	return node
}
