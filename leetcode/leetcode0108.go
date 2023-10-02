package leetcode
/*
面试题 01.08. 零矩阵
编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。
链接：https://leetcode.cn/problems/zero-matrix-lcci/
*/
func setZeroes(matrix [][]int)  {
	row:=make([]int,0)
	col:=make([]int,0)
	for i,v:=range matrix{
		for j,val:=range v{
			if val==0{
				row = append(row,i)
				col = append(col,j)
			}
		}
	}
	for i:=0;i<len(row);i++{
		 for k:=0;k<len(matrix[0]);k++{
			 matrix[row[i]][k]=0
		 }
		 for  k:=0;k<len(matrix);k++{
			 matrix[k][col[i]]=0
		 }
	}
	return

}