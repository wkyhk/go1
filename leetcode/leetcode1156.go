package leetcode

func MaxRepOpt1(text string) int {
	cnt := map[byte]int{}
	for i, _ := range text {
		cnt[text[i]]++
	}
	lo := 0
	hi := 0
	maxByte := text[lo]
	ans := 0
	maxCnt := 1
	mp := map[byte]int{}
	find := func() (int, byte) {
		m := 1
		c := text[lo]
		for k, v := range mp {
			if v > m {
				c, m = k, v
			}
		}
		return m, c
	}
	for hi < len(text) {
		mp[text[hi]]++
		maxCnt = max(maxCnt, mp[text[hi]])
		if maxCnt == mp[text[hi]] {
			maxByte = text[hi]
		}
		for !(hi-lo+1-maxCnt == 1 && cnt[maxByte] > maxCnt || hi-lo+1-maxCnt == 0) { //不满足条件，左指针右移
			mp[text[lo]]--
			maxCnt, maxByte = find()
			lo += 1
		}
		ans = max(ans, hi-lo+1)
		hi += 1
	}
	return ans
}
