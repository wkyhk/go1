package TreeNode

/*
145. 二叉树的后序遍历
链接：https://leetcode.cn/problems/binary-tree-postorder-traversal/description/
*/
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(postorderTraversal(root.Left), append(postorderTraversal(root.Right), root.Val)...)
}
