/*
 1123. 最深叶节点的最近公共祖先

给你一个有根节点 root 的二叉树，返回它 最深的叶节点的最近公共祖先 。

回想一下：

叶节点 是二叉树中没有子节点的节点
树的根节点的 深度 为 0，如果某一节点的深度为 d，那它的子节点的深度就是 d+1
如果我们假定 A 是一组节点 S 的 最近公共祖先，S 中的每个节点都在以 A 为根节点的子树中，且 A 的深度达到此条件下可能的最大值。
链接：https://leetcode.cn/problems/lowest-common-ancestor-of-deepest-leaves/description/?envType=daily-question&envId=2023-10-12
*/
package TreeNode

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, lca := f(root)
	return lca
}

func f(root *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}

	d1, lca1 := f(root.Left)
	h2, lca2 := f(root.Right)

	if d1 > h2 {
		return d1 + 1, lca1
	}
	if d1 < h2 {
		return h2 + 1, lca2
	}
	return d1 + 1, root
}
