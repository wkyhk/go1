package TreeNode

/*
113. 路径总和 II
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
叶子节点 是指没有子节点的节点。
链接：https://leetcode.cn/problems/path-sum-ii
*/
func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
	var dfs func(node *TreeNode, nums []int)
	dfs = func(node *TreeNode, nums []int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			sum := node.Val
			for _, v := range nums {
				sum += v
			}
			if targetSum == sum {
				nums = append(nums, node.Val)
				ans = append(ans, nums)
			}
		}

		temp := make([]int, len(nums))
		copy(temp, nums)
		temp = append(temp, node.Val)
		dfs(node.Right, temp)
		dfs(node.Left, temp)
	}
	dfs(root, []int{})
	return
}
