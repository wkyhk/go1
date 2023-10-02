package TreeNode

/*
1302. 层数最深叶子节点的和
给你一棵二叉树的根节点 root ，请你返回 层数最深的叶子节点的和 。
*/
func DeepestLeavesSum(root *TreeNode) int {
	maxlevel := -1
	sum := 0
	var dfs1 func(*TreeNode, int)
	dfs1 = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level > maxlevel {
			maxlevel = level
			sum = root.Val
		} else if level == maxlevel {
			sum += root.Val
		}
		dfs1(root.Left, level+1)
		dfs1(root.Right, level+1)
	}
	dfs1(root, 0)
	return sum

}
