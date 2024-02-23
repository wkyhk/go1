package TreeNode

/*
889. 根据前序和后序遍历构造二叉树
给定两个整数数组，preorder 和 postorder ，其中 preorder 是一个具有 无重复 值的二叉树的前序遍历，postorder 是同一棵树的后序遍历，重构并返回二叉树。
如果存在多个答案，您可以返回其中任何一个。
链接：https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal
*/
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	postMap := map[int]int{}
	for i, v := range postorder {
		postMap[v] = i
	}
	var dfs func(int, int, int, int) *TreeNode
	dfs = func(preLeft, preRight, postLeft, postRight int) *TreeNode {
		if preLeft > preRight {
			return nil
		}
		leftCount := 0
		if preLeft < preRight {
			leftCount = postMap[preorder[preLeft+1]] - postLeft + 1
		}
		return &TreeNode{
			Val:   preorder[preLeft],
			Left:  dfs(preLeft+1, preLeft+leftCount, postLeft, postLeft+leftCount-1),
			Right: dfs(preLeft+leftCount+1, preRight, postLeft+leftCount, postRight-1),
		}
	}
	return dfs(0, len(preorder)-1, 0, len(postorder)-1)
}
