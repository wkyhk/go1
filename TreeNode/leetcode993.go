package TreeNode

/*
993. 二叉树的堂兄弟节点
在二叉树中，根节点位于深度 0 处，每个深度为 k 的节点的子节点位于深度 k+1 处。
如果二叉树的两个节点深度相同，但 父节点不同 ，则它们是一对堂兄弟节点。
我们给出了具有唯一值的二叉树的根节点 root ，以及树中两个不同节点的值 x 和 y 。
只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true 。否则，返回 false。
链接：https://leetcode.cn/problems/cousins-in-binary-tree
*/
func isCousins(root *TreeNode, x int, y int) bool {
	xlevel, ylevel := -1, -1
	m := make(map[int]*TreeNode)
	var bfs func(*TreeNode, *TreeNode, int)
	bfs = func(tn, fn *TreeNode, i int) {
		if tn == nil {
			return
		}
		if tn.Val == x {
			xlevel = i
			m[x] = fn
		}
		if tn.Val == y {
			ylevel = i
			m[y] = fn
		}
		bfs(tn.Left, tn, i+1)
		bfs(tn.Right, tn, i+1)
	}
	bfs(root, nil, 0)
	if xlevel == ylevel && m[x] != m[y] {
		return true
	}
	return false
}
