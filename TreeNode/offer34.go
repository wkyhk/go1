package TreeNode
/*
剑指 Offer 34. 二叉树中和为某一值的路径
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
叶子节点 是指没有子节点的节点。
链接：https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum1(root *TreeNode, target int)(ans [][]int){
	var dfs func( *TreeNode,[]int,int)
	dfs = func (tree *TreeNode,nums []int,sum int){
		 if tree ==nil{
			 return
		 }
		 n:=make([]int,len(nums))
		 copy(n,nums)
		 if tree.Right!=nil||tree.Left!=nil{
			 n=append(n,tree.Val)
			 sum+=tree.Val
			 dfs(tree.Right,n,sum)
			 dfs(tree.Left,n,sum)
		 }else{
			 if sum+tree.Val==target{
				 n=append(n,tree.Val)
				 ans = append(ans,n)
			 }
		 }

	}
	dfs(root,[]int{},0)
    return
}