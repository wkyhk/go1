package main

import "fmt"

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	ans := 0
	if k > numOnes {
		ans = numOnes
		k -= numOnes
	} else {
		return k
	}
	if k > numZeros {
		k -= numZeros
	} else {
		return ans
	}

	return ans - k
}

func main() {

	fmt.Println(kItemsWithMaximumSum(3, 2, 0, 2))

}

/* func main() {
	start := time.Now()
	var wg sync.WaitGroup
	ans := make([]int, 0)
	for i := 1; i < 65535; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			add := fmt.Sprintf("localhost:%d", j)
			conn, err := net.Dial("tcp", add)

			if err != nil {
				return
			}
			conn.Close()
			ans = append(ans, j)
		}(i)
	}
	wg.Wait()
	end := time.Since(start) / 1e9
	fmt.Println(ans)
	fmt.Printf("%d seconds", end)
}
*/
