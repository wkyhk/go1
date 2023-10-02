package leetcode

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func PruneTree(root *TreeNode) *TreeNode {
	if root == nil {

		return root
	}
	root.Left = PruneTree(root.Left)
	root.Right = PruneTree(root.Right)
	if root.Right == nil && root.Left == nil && root.Val == 0 {
		root = nil

	}
	return root
}
