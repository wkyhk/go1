package leetcode

/*
1491. 去掉最低工资和最高工资后的工资平均值
给你一个整数数组 salary ，数组里每个数都是 唯一 的，其中 salary[i] 是第 i 个员工的工资。

请你返回去掉最低工资和最高工资以后，剩下员工工资的平均值。
链接：https://leetcode.cn/problems/average-salary-excluding-the-minimum-and-maximum-salary
*/
func average(salary []int) float64 {
	minSal, maxSal := salary[0], salary[0]
	sum := 0
	for _, v := range salary {
		sum += v
		if minSal > v {
			minSal = v
		}
		if maxSal < v {
			maxSal = v
		}
	}
	return float64(sum-minSal-maxSal) / float64(len(salary)-2)
}
