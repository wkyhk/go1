package offer

import (
     "fmt"
     "math"
     "sort"
)

/*
剑指 Offer II 008. 和大于等于 target 的最短子数组
给定一个含有n个正整数的数组和一个正整数 target 。
找出该数组中满足其和 ≥ target 的长度最小的 连续子数组[numsl, numsl+1, ..., numsr-1, numsr] ，
并返回其长度。如果不存在符合条件的子数组，返回 0 。
链接：https://leetcode.cn/problems/2VG8Kg
*/
func minSubArrayLen(s int, nums []int) int {
     n := len(nums)
     if n == 0 {
          return 0
     }
     ans := math.MaxInt32
     sums := make([]int, n + 1)
     // 为了方便计算，令 size = n + 1
     // sums[0] = 0 意味着前 0 个元素的前缀和为 0
     // sums[1] = A[0] 前 1 个元素的前缀和为 A[0]
     // 以此类推
     for i := 1; i <= n; i++ {
          sums[i] = sums[i - 1] + nums[i - 1]
     }
     for i := 1; i <= n; i++ {
          target := s + sums[i-1]
          bound := sort.SearchInts(sums, target)
          if bound < 0 {
               bound = -bound - 1
          }
          if bound <= n {
               ans = min(ans, bound - (i - 1))
          }
     }
     if ans == math.MaxInt32 {
          return 0
     }
     return ans
}

func min(x, y int) int {
     if x < y {
          return x
     }
     return y
}


func TestOffer08(){
     fmt.Println(7,[]int{2,3,1,2,4,3})
     fmt.Println(4,[]int{1,4,4})
     fmt.Println(11,[]int{1,1,1,1,1,1,1,1})
}