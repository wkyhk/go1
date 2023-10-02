package TreeNode

import "fmt"

/*
144. 二叉树的前序遍历
给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
https://leetcode.cn/problems/binary-tree-preorder-traversal/
*/
func preorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	var pro func(*TreeNode)
	pro = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		pro(node.Left)
		pro(node.Right)
	}
	pro(root)
	return ans
}
func Example114() {
	root := New()
	fmt.Println(preorderTraversal(root))
}
