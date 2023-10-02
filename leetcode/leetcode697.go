package leetcode

/*
697. 数组的度
给定一个非空且只包含非负数的整数数组nums，数组的 度 的定义是指数组里任一元素出现频数的最大值。
你的任务是在 nums 中找到与nums拥有相同大小的度的最短连续子数组，返回其长度。
链接：https://leetcode.cn/problems/degree-of-an-array
*/
type entry struct{ cnt, l, r int }

func findShortestSubArray(nums []int) (ans int) {
	mp := map[int]entry{}
	for i, v := range nums {
		if e, has := mp[v]; has {
			e.cnt++
			e.r = i
			mp[v] = e
		} else {
			mp[v] = entry{1, i, i}
		}
	}
	maxCnt := 0
	for _, e := range mp {
		if e.cnt > maxCnt {
			maxCnt, ans = e.cnt, e.r-e.l+1
		} else if e.cnt == maxCnt {
			ans = min(ans, e.r-e.l+1)
		}
	}
	return
}

/*func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}*/
