package TreeNode

/*
1161. 最大层内元素和
给你一个二叉树的根节点 root。设根节点位于二叉树的第 1 层，而根节点的子节点位于第 2 层，依此类推。
请返回层内元素之和 最大 的那几层（可能只有一层）的层号，并返回其中 最小 的那个。
链接：https://leetcode.cn/problems/maximum-level-sum-of-a-binary-tree
树中的节点数在 [1, 104]范围内
-105 <= Node.val <= 105
*/
func MaxLevelSum(root *TreeNode) int {
	sum := make([]int, 0)
	dfs(root, 0, &sum)
	ans := 0
	for i, v := range sum {
		if v > sum[ans] {
			ans = i
		}
	}
	return ans + 1
}
func dfs(root *TreeNode, level int, sum *[]int) {
	if root == nil {
		return
	}
	if level == len(*sum) {
		*sum = append(*sum, root.Val)
	} else {
		(*sum)[level] += root.Val
	}
	if root.Right != nil {
		dfs(root.Right, level+1, sum)
	}
	if root.Left != nil {
		dfs(root.Left, level+1, sum)
	}

}
