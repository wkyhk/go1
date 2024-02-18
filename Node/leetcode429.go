package Node

/*
429.N叉树的层序遍历
给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。

树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）
链接：https://leetcode.cn/problems/n-ary-tree-level-order-traversal/description/
*/
func levelOrder(root *Node) (res [][]int) {
	if root == nil {
		return
	}
	var dfs func(*Node, int)
	dfs = func(node *Node, level int) {
		if node == nil {
			return
		}
		if len(res) == level {
			res = append(res, []int{})

		}
		res[level] = append(res[level], node.Val)
		for _, v := range node.Children {
			dfs(v, level+1)
		}
	}
	dfs(root, 0)
	return
}
