package TreeNode

/*
637. 二叉树的层平均值
给定一个非空二叉树的根节点 root , 以数组的形式返回每一层节点的平均值。与实际答案相差 10-5 以内的答案可以被接受。
链接：https://leetcode.cn/problems/average-of-levels-in-binary-tree/
*/
func averageOfLevels(root *TreeNode) []float64 {
	ans := make([]float64, 0)
	count := make([]int, 0)
	var bfs func(tree *TreeNode, level int)
	bfs = func(tree *TreeNode, level int) {
		if tree == nil {
			return
		}
		if level >= len(ans) {
			ans = append(ans, float64(tree.Val))
			count = append(count, 1)
		} else {
			ans[level] += float64(tree.Val)
			count[level]++
		}
		bfs(tree.Left, level+1)
		bfs(tree.Right, level+1)
	}
	bfs(root, 0)
	for i, _ := range ans {
		ans[i] = ans[i] / float64(count[i])
	}
	return ans
}
