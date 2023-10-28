package main

import "GO1/leetcode"

func main() {
	leetcode.TestmamatrixSum()
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
