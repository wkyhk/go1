package TreeNode

/*
559. N 叉树的最大深度
给定一个 N 叉树，找到其最大深度。
最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。
N 叉树输入按层序遍历序列化表示，每组子节点由空值分隔（请参见示例）。
链接：https://leetcode.cn/problems/maximum-depth-of-n-ary-tree
*/
func maxDepth1(root *Node) int {
	deep := 0
	var dfs func(*Node, int)
	dfs = func(node *Node, level int) {
		if node == nil {
			return
		}
		level = level + 1
		deep = max(deep, level)
		for i := 0; i < len(node.Children); i++ {
			dfs(node.Children[i], level)
		}
	}
	dfs(root, deep)
	return deep
}
