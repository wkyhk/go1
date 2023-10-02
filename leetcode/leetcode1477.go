package leetcode

import "math"

/*
1477. 找两个和为目标值且不重叠的子数组
给你一个整数数组arr 和一个整数值target。
请你在 arr中找 两个互不重叠的子数组且它们的和都等于target。可能会有多种方案，请你返回满足要求的两个子数组长度和的 最小值 。
请返回满足要求的最小长度和，如果无法找到这样的两个子数组，请返回 -1
链接：https://leetcode.cn/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum
*/
func minSumOfLengths(arr []int, target int) int {

	if len(arr) <= 1 {
		return -1
	}
	result := comResult(arr, target)
	if len(result) < 2 {
		return -1
	}
	pre := comPre(arr, result)
	suf := comSuf(arr, result)

	m := math.MaxInt64
	for i := 0; i < len(arr)-1; i++ {
		if pre[i] == math.MaxInt64 || suf[i+1] == math.MaxInt64 {
			continue
		}
		if pre[i]+suf[i+1] < m {
			m = pre[i] + suf[i+1]
		}
	}
	if m == math.MaxInt64 {
		return -1
	}
	return m
}

func comResult(arr []int, target int) [][3]int {
	la := len(arr)
	result := make([][3]int, 0, 10)
	lo, hi := 0, 0
	sum := arr[0]
	for hi < la {
		if sum < target {
			hi++
			if hi < la {
				sum += arr[hi]
			}
			continue
		}
		if sum > target {
			sum -= arr[lo]
			if lo == hi {
				hi++
				if hi < la {
					sum += arr[hi]
				}
			}
			lo++
			continue
		}
		result = append(result, [3]int{lo, hi, hi - lo + 1})
		hi++
		if hi < la {
			sum += arr[hi]
		}
	}
	return result
}
func comPre(arr []int, result [][3]int) []int {
	ret := make([]int, len(arr))
	pre := math.MaxInt64
	var reti, ri = 0, 0
	for i := 0; i < len(arr); i++ {
		if ri >= len(result) {
			break
		}
		for ; reti < result[ri][1]; reti++ {
			ret[reti] = pre
		}
		if result[ri][2] < pre {
			ret[reti] = result[ri][2]
			pre = ret[reti]
		}
		ri++
	}
	for i := 1; i < len(arr); i++ {
		if ret[i] == 0 {
			ret[i] = ret[i-1]
		}
	}
	return ret
}

func comSuf(arr []int, result [][3]int) []int {
	ret := make([]int, len(arr))
	suf := math.MaxInt64
	var reti, ri = len(arr) - 1, len(result) - 1
	for i := len(arr) - 1; i >= 0; i-- {
		if ri < 0 {
			break
		}
		for ; reti > result[ri][0]; reti-- {
			ret[reti] = suf
		}
		if result[ri][2] < suf {
			ret[reti] = result[ri][2]
			suf = ret[reti]
		}
		ri--
	}
	for i := len(arr) - 2; i >= 0; i-- {
		if ret[i] == 0 {
			ret[i] = ret[i+1]
		}
	}
	return ret
}
