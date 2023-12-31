package leetcode

import (
	"strconv"
	"strings"
)

/*
1154. 一年中的第几天
给你一个字符串 date ，按 YYYY-MM-DD 格式表示一个 现行公元纪年法 日期。返回该日期是当年的第几天。
链接：https://leetcode.cn/problems/day-of-the-year/
*/
func dayOfYear(date string) int {
	str := strings.Split(date, "-")
	year, _ := strconv.Atoi(str[0])
	month, _ := strconv.Atoi(str[1])
	day, _ := strconv.Atoi(str[2])
	res := 0
	for i := 1; i < month; i++ {
		if i == 1 || i == 3 || i == 5 || i == 7 || i == 8 || i == 10 || i == 12 {
			res += 31
		} else if i == 4 || i == 6 || i == 9 || i == 11 {
			res += 30
		} else {
			if year%4 == 0 && year%100 != 0 || year%400 == 0 {
				res += 29
			} else {
				res += 28
			}
		}

	}
	return res + day
}
