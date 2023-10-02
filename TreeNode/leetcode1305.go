package TreeNode

import (
	"sort"
)

/*
1305. 两棵二叉搜索树中的所有元素
给你 root1 和 root2 这两棵二叉搜索树。请你返回一个列表，其中包含 两棵树 中的所有整数并按 升序 排序。.
*/
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	ans := []int{}
	var dfs1 func(*TreeNode)
	dfs1 = func(root *TreeNode) {
		if root == nil {
			return
		}
		ans = append(ans, root.Val)
		dfs1(root.Left)
		dfs1(root.Right)
	}
	dfs1(root1)
	dfs1(root2)
	sort.Ints(ans)
	return ans

}
