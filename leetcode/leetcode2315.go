package leetcode

import "fmt"

/*
2315. 统计星号
给你一个字符串 s ，每 两个 连续竖线 '|' 为 一对 。换言之，第一个和第二个 '|' 为一对，第三个和第四个 '|' 为一对，以此类推。
请你返回 不在 竖线对之间，s 中 '*' 的数目。
注意，每个竖线 '|' 都会 恰好 属于一个对。
链接：https://leetcode.cn/problems/count-asterisks
*/
func countAsterisks(s string) int {
	ans := 0
	is := true
	for _, v := range []byte(s) {
		if is {
			if v == '*' {
				ans++
			}
		}
		if v == '|' {
			is = !is
		}
	}
	return ans
}
func TestLeetCode2315() {
	fmt.Println(countAsterisks("l|*e*et|c**o|*de|"))
	fmt.Println(countAsterisks("iamprogrammer"))
	fmt.Println(countAsterisks("yo|uar|e**|b|e***au|tifu|l"))

}
