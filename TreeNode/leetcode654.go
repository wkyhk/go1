package TreeNode

/*
654. 最大二叉树
给定一个不重复的整数数组 nums 。 最大二叉树 可以用下面的算法从 nums 递归地构建:创建一个根节点，其值为 nums 中的最大值。
递归地在最大值 左边 的 子数组前缀上 构建左子树。
递归地在最大值 右边 的 子数组后缀上 构建右子树。
返回 nums 构建的 最大二叉树 。
链接：https://leetcode.cn/problems/maximum-binary-tree
*/
func constructMaximumBinaryTree(nums []int) *TreeNode {
	tree := make([]*TreeNode, len(nums))
	stk := []int{}
	for i, num := range nums {
		tree[i] = &TreeNode{Val: num}
		for len(stk) > 0 && num > nums[stk[len(stk)-1]] {
			tree[i].Left = tree[stk[len(stk)-1]]
			stk = stk[:len(stk)-1]
		}
		if len(stk) > 0 {
			tree[stk[len(stk)-1]].Right = tree[i]
		}
		stk = append(stk, i)
	}
	return tree[stk[0]]
}
