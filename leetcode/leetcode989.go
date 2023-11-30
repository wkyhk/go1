package leetcode

/*
989. 数组形式的整数加法
整数的 数组形式  num 是按照从左到右的顺序表示其数字的数组。
例如，对于 num = 1321 ，数组形式是 [1,3,2,1] 。
给定 num ，整数的 数组形式 ，和整数 k ，返回 整数 num + k 的 数组形式 。
链接：https://leetcode.cn/problems/add-to-array-form-of-integer/description/
*/
func addToArrayForm(num []int, k int) (ans []int) {
	for i := len(num) - 1; i >= 0; i-- {
		sum := num[i] + k%10
		k /= 10
		if sum >= 10 {
			k++
			sum -= 10
		}
		ans = append(ans, sum)
	}
	for ; k > 0; k /= 10 {
		ans = append(ans, k%10)
	}
	reverse1(ans)
	return
}

func reverse1(num []int) {
	for i, n := 0, len(num); i < n/2; i++ {
		num[i], num[n-1-i] = num[n-1-i], num[i]
	}
}
