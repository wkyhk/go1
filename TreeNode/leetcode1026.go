package TreeNode

/*
1026. 节点与其祖先之间的最大差值
给定二叉树的根节点root，找出存在于 不同 节点A 和B之间的最大值 V，其中V = |A.val - B.val|，且A是B的祖先。
（如果 A 的任何子节点之一为 B，或者 A 的任何子节点是 B 的祖先，那么我们认为 A 是 B 的祖先）
https://leetcode.cn/problems/maximum-difference-between-node-and-ancestor/
*/
func maxAncestorDiff(root *TreeNode) int {
	ans := 0
	var bfs func(*TreeNode, int, int)
	bfs = func(node *TreeNode, ma, mi int) {
		if node == nil {
			ans = max1(ans, ma-mi)
			return
		}
		if node.Val > ma {
			ma = node.Val
		}
		if node.Val < mi {
			mi = node.Val
		}
		bfs(node.Left, ma, mi)
		bfs(node.Right, ma, mi)
	}
	bfs(root, root.Val, root.Val)
	return ans
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
