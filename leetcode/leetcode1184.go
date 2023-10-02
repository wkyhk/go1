package leetcode

func DistanceBetweenBusStops(distance []int, start, destination int) int {
	if start > destination {
		start, destination = destination, start
	}
	sum1, sum2 := 0, 0
	for _, v := range distance {
		sum2 += v
	}

	for i := start; i < destination; i++ {
		sum1 += distance[i]
	}
	return min(sum1, sum2-sum1)
}
