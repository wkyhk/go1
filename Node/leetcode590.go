package Node

/*
590. N 叉树的后序遍历
给定一个 n 叉树的根节点 root ，返回 其节点值的 后序遍历 。

n 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。
链接：https://leetcode.cn/problems/n-ary-tree-postorder-traversal/description/
*/
func postorder(root *Node) []int {
	res := make([]int, 0)
	var BFS func(root *Node)
	BFS = func(root *Node) {
		if root == nil {
			return
		}
		for _, v := range root.Children {
			BFS(v)

		}
		res = append(res, root.Val)
	}
	BFS(root)
	return res
}
