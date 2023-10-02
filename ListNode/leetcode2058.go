package listnode

import "math"

/*
2058. 找出临界点之间的最小和最大距离
链表中的 临界点 定义为一个 局部极大值点 或 局部极小值点 。
如果当前节点的值 严格大于 前一个节点和后一个节点，那么这个节点就是一个  局部极大值点 。
如果当前节点的值 严格小于 前一个节点和后一个节点，那么这个节点就是一个  局部极小值点 。
注意：节点只有在同时存在前一个节点和后一个节点的情况下，才能成为一个 局部极大值点 / 极小值点 。
给你一个链表 head ，返回一个长度为 2 的数组 [minDistance, maxDistance] ，其中 minDistance 是任意两个不同临界点之间的最
小距离，maxDistance 是任意两个不同临界点之间的最大距离。如果临界点少于两个，则返回 [-1，-1]
链接：https://leetcode.cn/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points
*/
func nodesBetweenCriticalPoints(head *ListNode) []int {
	a, b, c := head, head.Next, head.Next.Next
	first, last, minDis := 0, 0, math.MaxInt32
	for i, prev := 1, 0; c != nil; i++ { // 遍历链表，寻找临界点
		if a.Val < b.Val && b.Val > c.Val || a.Val > b.Val && b.Val < c.Val {
			if first == 0 {
				first = i // 首个临界点位置
			}
			last = i // 最末临界点位置
			if prev > 0 && i-prev < minDis {
				minDis = i - prev // 更新相邻临界点位置之差的最小值
			}
			prev = i
		}
		a, b, c = b, c, c.Next
	}
	if first == last { // 临界点少于两个
		return []int{-1, -1}
	}
	return []int{minDis, last - first}
}
