package main

import "fmt"

func solve(n int, nums []int) int {
	set := make(map[int]int)
	maxLen := 0
	l := 0

	for r, num := range nums {
		if prev, exists := set[num]; exists && prev >= l {
			l = prev + 1
		}

		set[num] = r
		maxLen = max(maxLen, r-l+1)
	}

	return max(maxLen, n-l)
}

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Println(solve(n, nums))
}
