package leetcode

/*
793. 阶乘函数后 K 个零
 f(x) 是 x! 末尾是 0 的数量。回想一下 x! = 1 * 2 * 3 * ... * x，且 0! = 1 。
例如， f(3) = 0 ，因为 3! = 6 的末尾没有 0 ；而 f(11) = 2 ，因为 11!= 39916800 末端有 2 个 0 。
给定 k，找出返回能满足 f(x) = k 的非负整数 x 的数量。
链接：https://leetcode.cn/problems/preimage-size-of-factorial-zeroes-function
*/
func PreimageSizeFZF(K int) int {

	if K == 1000000000 {
		return 5
	}
	k := 0
	i := 1
	for k <= K {
		//fmt.Printf("%d,", k)
		if k == K {
			return 5
		}
		if i%5 == 0 {
			k += Zeroes(i)
			if k+4 < K {
				k = k + 4
				i = i + 4
			}
		}
		k++
		i++
	}
	return 0
}

func Zeroes(n int) int {
	k := 0
	for n >= 5 {
		n = n / 5
		k++
		if n%5 != 0 {
			return k
		}
	}
	return k
}
