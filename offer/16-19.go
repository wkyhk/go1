/*
面试题 16.19. 水域大小
你有一个用于表示一片土地的整数矩阵land，该矩阵中每个点的值代表对应地点的海拔高度。若值为0则表示水域。由垂直、水平或对角连接的水域为池塘。池塘的大小是指相连接的水域的个数。编写一个方法来计算矩阵中所有池塘的大小，返回值需要从小到大排序。
链接：https://leetcode.cn/problems/pond-sizes-lcci/?envType=daily-question&envId=2023-10-12
*/
package offer

import "sort"

func pondSizes(land [][]int) []int {
	m, n := len(land), len(land[0])
	var dfs func(int, int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= m || y < 0 || y >= n || land[x][y] != 0 {
			return 0
		}
		land[x][y] = -1
		res := 1
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if dx == 0 && dy == 0 {
					continue
				}
				res += dfs(x+dx, y+dy)
			}
		}
		return res
	}
	res := []int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if land[i][j] == 0 {
				res = append(res, dfs(i, j))
			}
		}
	}
	sort.Ints(res)
	return res
}
