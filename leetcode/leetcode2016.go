package leetcode

import "fmt"

/*
2016. 增量元素之间的最大差值
给你一个下标从 0 开始的整数数组 nums ，该数组的大小为 n ，请你计算 nums[j] - nums[i] 能求得的 最大差值 ，
其中 0 <= i < j < n 且 nums[i] < nums[j] 。
返回 最大差值 。如果不存在满足要求的 i 和 j ，返回 -1 。
链接：https://leetcode.cn/problems/maximum-difference-between-increasing-elements
*/
func maximumDifference(nums []int) int {
  ans:=-1
  min := -1
  for i:=0;i<len(nums);i++{
    if i==0{
      min = nums[i]
    } else if i>0 {
      if min > nums[i] {
        min = nums[i]
      } else if min < nums[i] {
        if ans < nums[i]-min {
          ans = nums[i] - min
        }
      }
    }
  }
  return  ans
}
func TestLeetCode2016(){
  fmt.Println(maximumDifference([]int{7,1,5,4}))
  fmt.Println(maximumDifference([]int{9,4,3,2}))
  fmt.Println(maximumDifference([]int{1,5,2,10}))
}