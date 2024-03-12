package TreeNode

/*
1261. 在受污染的二叉树中查找元素
给出一个满足下述规则的二叉树：

root.val == 0
如果 treeNode.val == x 且 treeNode.left != null，那么 treeNode.left.val == 2 * x + 1
如果 treeNode.val == x 且 treeNode.right != null，那么 treeNode.right.val == 2 * x + 2
现在这个二叉树受到「污染」，所有的 treeNode.val 都变成了 -1。

请你先还原二叉树，然后实现 FindElements 类：

FindElements(TreeNode* root) 用受污染的二叉树初始化对象，你需要先把它还原。
bool find(int target) 判断目标值 target 是否存在于还原后的二叉树中并返回结果。
链接：https://leetcode.cn/problems/find-elements-in-a-contaminated-binary-tree/description/
*/

type FindElements struct {
	Val map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	if root == nil {
		return FindElements{}
	}
	valmap := make(map[int]bool)
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, val int) {
		if root == nil {
			return
		}
		root.Val = val
		valmap[val] = true
		dfs(root.Left, 2*val+1)
		dfs(root.Right, 2*val+2)

	}
	dfs(root, 0)
	return FindElements{Val: valmap}
}

func (this *FindElements) Find(target int) bool {
	return this.Val[target]
}
