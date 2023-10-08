/*
 771. 宝石与石头
    给你一个字符串 jewels 代表石头中宝石的类型，另有一个字符串 stones 代表你拥有的石头。 stones 中每个字符代表了一种你拥有的石头的类型，你想知道你拥有的石头中有多少是宝石。

字母区分大小写，因此 "a" 和 "A" 是不同类型的石头。
链接：https://leetcode.cn/problems/jewels-and-stones/?envType=daily-question&envId=2023-10-08
*/
package leetcode

func numJewelsInStones(jewels string, stones string) int {
	var count int
	m := make(map[byte]struct{}, 0)
	for i := 0; i < len(jewels); i++ {
		m[jewels[i]] = struct{}{}
	}
	for i := 0; i < len(stones); i++ {
		if _, ok := m[stones[i]]; ok {
			count++
		}
	}
	return count
}
