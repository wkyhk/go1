package leetcode

/*
1828. 统计一个圆中点的数目
给你一个数组points，其中points[i] = [xi, yi]，表示第i个点在二维平面上的坐标。多个点可能会有 相同的坐标。
同时给你一个数组queries，其中queries[j] = [xj, yj, rj]，表示一个圆心在(xj, yj)且半径为rj的圆。
对于每一个查询queries[j]，计算在第 j个圆 内点的数目。如果一个点在圆的 边界上，我们同样认为它在圆内。
请你返回一个数组answer，其中answer[j]是第j个查询的答案。
链接：https://leetcode.cn/problems/queries-on-number-of-points-inside-a-circle
*/
func countPoints(points [][]int, queries [][]int) []int {
     ans := make([]int, len(queries))
     for i, q := range queries {
          x, y, r := q[0], q[1], q[2]
          for _, p := range points {
               if (p[0]-x)*(p[0]-x)+(p[1]-y)*(p[1]-y) <= r*r {
                    ans[i]++
               }
          }
     }
     return ans
}

