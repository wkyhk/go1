package TreeNode

/*
700. 二叉搜索树中的搜索
给定二叉搜索树（BST）的根节点 root 和一个整数值 val。
你需要在 BST 中找到节点值等于 val 的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 null 。
链接：https://leetcode.cn/problems/search-in-a-binary-search-tree
*/
func searchBST(root *TreeNode, val int) *TreeNode {
	var ans *TreeNode = nil
	for root != nil {
		if root.Val == val {
			ans = root
			break
		} else if root.Val > val {
			root = root.Left
		} else if root.Val < val {
			root = root.Right
		}
	}
	return ans
}
