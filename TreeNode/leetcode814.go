package TreeNode

/*
814. 二叉树剪枝
给你二叉树的根结点 root ，此外树的每个结点的值要么是 0 ，要么是 1 。
返回移除了所有不包含 1 的子树的原二叉树。
节点 node 的子树为 node 本身加上所有 node 的后代。
链接：https://leetcode.cn/problems/binary-tree-pruning
*/
func pruneTree(root *TreeNode) *TreeNode {
	ans := root
	var bfs func(*TreeNode) bool
	bfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		a := bfs(node.Right)
		b := bfs(node.Left)
		if a {
			node.Right = nil
		}
		if b {
			node.Left = nil
		}
		if node.Val == 0 {
			if a && b {
				return true
			}
		}
		return false
	}
	if bfs(root) {
		return ans
	}
	return nil
}
