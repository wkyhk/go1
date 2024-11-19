package leetcode

import "math"

/*
2708. 一个小组的最大实力值
给你一个下标从 0 开始的整数数组 nums ，它表示一个班级中所有学生在一次考试中的成绩。老师想选出一部分同学组成一个 非空 小组，且这个小组的 实力值 最大，
如果这个小组里的学生下标为 i0, i1, i2, ... , ik ，那么这个小组的实力值定义为 nums[i0] * nums[i1] * nums[i2] * ... * nums[ik] 。
请你返回老师创建的小组能得到的最大实力值为多少。
链接：https://leetcode.cn/problems/maximum-strength-of-a-group/description/
*/
func maxStrength(nums []int) int64 {
	negativeCount, zeroCount, positiveCount := 0, 0, 0
	prod := 1
	maxNegative := math.MinInt32

	for _, num := range nums {
		if num < 0 {
			negativeCount++
			prod *= num
			if num > maxNegative {
				maxNegative = num
			}
		} else if num == 0 {
			zeroCount++
		} else {
			prod *= num
			positiveCount++
		}
	}

	if negativeCount == 1 && zeroCount == 0 && positiveCount == 0 {
		return int64(nums[0])
	}
	if negativeCount <= 1 && positiveCount == 0 {
		return int64(0)
	}
	if prod < 0 {
		return int64(prod / maxNegative)
	} else {
		return int64(prod)
	}
}
