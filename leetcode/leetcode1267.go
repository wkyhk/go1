package leetcode

/*
1267. 统计参与通信的服务器
这里有一幅服务器分布图，服务器的位置标识在 m * n 的整数矩阵网格 grid 中，1 表示单元格上有服务器，0 表示没有。
如果两台服务器位于同一行或者同一列，我们就认为它们之间可以进行通信。
请你统计并返回能够与至少一台其他服务器进行通信的服务器的数量。
链接：https://leetcode.cn/problems/count-servers-that-communicate
*/
func CountServers(grid [][]int) int {
	count := 0
	for i, v := range grid {
		for j, v1 := range v {
			if v1 == 1 {
				b := false
				for i1 := 0; i1 < len(grid); i1++ {
					if grid[i1][j] == 1 && i1 != i {
						count++
						b = true
						break
					}
				}
				if !b {
					for j1 := 0; j1 < len(grid[0]); j1++ {
						if grid[i][j1] == 1 && j1 != j {
							count++
							break
						}
					}
				}
			}
		}
	}
	return count
}
