package leetcode

/*
1496. 判断路径是否相交
给你一个字符串 path，其中 path[i] 的值可以是 'N'、'S'、'E' 或者 'W'，分别表示向北、向南、向东、向西移动一个单位。

你从二维平面上的原点 (0, 0) 处开始出发，按 path 所指示的路径行走。

如果路径在任何位置上与自身相交，也就是走到之前已经走过的位置，请返回 true ；否则，返回 false 。



来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/path-crossing
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func getPoint(x, y int) int {
	return x*20001 + y
}
func IsPathCrossing(path string) bool {
	mp := make(map[int]struct{}, 0)
	x, y := 0, 0
	mp[getPoint(x, y)] = struct{}{}
	for _, v := range path {
		switch v {
		case 'N':
			x++
		case 'S':
			x--
		case 'E':
			y++
		case 'W':
			y--
		}
		pointer := getPoint(x, y)
		if _, ok := mp[pointer]; ok {
			return true
		} else {
			mp[pointer] = struct{}{}
		}
	}
	return false
}
