package leetcode

/*
1010. 总持续时间可被 60 整除的歌曲
在歌曲列表中，第 i 首歌曲的持续时间为 time[i] 秒。
返回其总持续时间（以秒为单位）可被 60 整除的歌曲对的数量。形式上，我们希望下标数字 i 和 j 满足  i < j 且有 (time[i] + time[j]) % 60 == 0。
链接：https://leetcode.cn/problems/pairs-of-songs-with-total-durations-divisible-by-60
*/
//直接相加，会超时
/*  func numPairsDivisibleBy60(time []int) int {
	  ans:=0
      for i:=0;i<len(time)-1;i++{
		for j:=i+1;j<len(time);j++{
			if (time[i]+time[j])%60==0{
				ans++
			}
		}
	  }
	  return ans
} */
func numPairsDivisibleBy60(time []int) int {
	m := make([]int, 60)
	cnt := 0
	for _, n := range time {
		if n%60 == 0 {
			cnt += m[0]
		} else {
			cnt += m[60-n%60]
		}
		m[n%60]++
	}
	return cnt

}
