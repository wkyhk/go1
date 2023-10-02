package leetcode
/*
2319：判断一个矩阵是否为X矩阵
如果一个正方形矩阵满足下述 全部 条件，则称之为一个 X 矩阵 ：
矩阵对角线上的所有元素都 不是 0
矩阵中所有其他元素都是 0
给你一个大小为 n x n 的二维整数数组 grid ，表示一个正方形矩阵。如果 grid 是一个 X 矩阵 ，返回 true ；
否则，返回 false 。
链接：https://leetcode.cn/problems/check-if-matrix-is-x-matrix
*/
func checkXMatrix(grid [][]int) bool {
    for i,v:=range grid{
		for j,v1:=range v{
			if i==j||i+j==len(grid)-1{
				if v1==0{
					return false
				}
			}else{
				if v1!=0{
					return false
				}
			}
		}
	}
	return true
}