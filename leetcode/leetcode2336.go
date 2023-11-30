package leetcode

/*
2336. 无限集中的最小数字
现有一个包含所有正整数的集合 [1, 2, 3, 4, 5, ...] 。

实现 SmallestInfiniteSet 类：

SmallestInfiniteSet() 初始化 SmallestInfiniteSet 对象以包含 所有 正整数。
int popSmallest() 移除 并返回该无限集中的最小整数。
void addBack(int num) 如果正整数 num 不 存在于无限集中，则将一个 num 添加 到该无限集中。
链接：https://leetcode.cn/problems/smallest-number-in-infinite-set/
*/

type SmallestInfiniteSet struct {
	nums [1000]int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	res := 0
	for i, v := range s.nums {
		if v == 0 {
			res = i + 1
			s.nums[i] = 1
			break
		}
	}
	return res
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	if s.nums[num-1] == 1 {
		s.nums[num-1] = 0
	}
}
