package TreeNode

/*
94. 二叉树的中序遍历
链接：https://leetcode.cn/problems/binary-tree-inorder-traversal/description/
*/
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(inorderTraversal(root.Left), append([]int{root.Val}, inorderTraversal(root.Right)...)...)
}
