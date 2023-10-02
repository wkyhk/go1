package leetcode

/*
448. 找到所有数组中消失的数字
给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，
并以数组的形式返回结果。
链接：https://leetcode.cn/problems/find-all-numbers-disappeared-in-an-array
*/
func findDisappearedNumbers(nums []int) []int {
	m := make(map[int]struct{}, 0)
	ans := make([]int, 0)
	for _, v := range nums {
		m[v] = struct{}{}
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			ans = append(ans, i)
		}
	}
	return ans
}

/*//不使用额外空间
func findDisappearedNumbers1(nums []int) []int {
     sort.Ints(nums)
     ans:=make([]int,0)
     for i,v:=range nums{
          if i+1>v{
               ans=append(ans,v)
          }
     }
     return  ans
}*/
