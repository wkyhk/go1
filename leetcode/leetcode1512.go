package leetcode

//1512. 好数对的数目
/*
   给你一个整数数组 nums 。

如果一组数字 (i,j) 满足 nums[i] == nums[j] 且 i < j ，就可以认为这是一组 好数对 。

返回好数对的数目。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/number-of-good-pairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func NumIdenticalPairs(nums []int) int {
	mp := make(map[int]int, 0)
	ans := 0
	for _, v := range nums {
		mp[v]++
	}
	for _, v := range mp {
		ans += (v - 1) * v / 2
	}
	return ans
}
