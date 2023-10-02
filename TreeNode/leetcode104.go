package TreeNode

/*
 104. 二叉树的最大深度
给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var dfs func(*TreeNode, int) int
	dfs = func(tn *TreeNode, i int) int {
		if tn == nil {
			return i - 1
		}
		if tn.Left == nil && tn.Right == nil {
			return i
		}
		return max(dfs(tn.Right, i+1), dfs(tn.Left, i+1))
	}
	return dfs(root, 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/* func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
*/
