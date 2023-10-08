/*
2437. 有效时间的数目
给你一个长度为 5 的字符串 time ，表示一个电子时钟当前的时间，格式为 "hh:mm" 。最早 可能的时间是 "00:00" ，最晚 可能的时间是 "23:59" 。
在字符串 time 中，被字符 ? 替换掉的数位是 未知的 ，被替换的数字可能是 0 到 9 中的任何一个。
请你返回一个整数 answer ，将每一个 ? 都用 0 到 9 中一个数字替换后，可以得到的有效时间的数目。
链接：https://leetcode.cn/problems/number-of-valid-clock-times/?envType=daily-question&envId=2023-10-05
*/
package leetcode

func countTime(time string) int {
	num := 1
	for i := 0; i < 5; i++ {
		if i == 0 && time[i] == '?' {
			if time[i+1] == '?' {
				num *= 24
			} else if time[i+1] >= '4' {
				num *= 2
			} else if time[i+1] < '4' {
				num *= 3
			}
		}
		if i == 1 && time[i] == '?' {
			if time[i-1] == '?' {
				num *= 1
			} else if time[i-1] > '1' {
				num *= 4
			} else if time[i-1] <= '1' {
				num *= 10
			}
		}
		if i == 3 && time[i] == '?' {
			num *= 6
		}
		if i == 4 && time[i] == '?' {
			num *= 10
		}
	}
	return num
}
