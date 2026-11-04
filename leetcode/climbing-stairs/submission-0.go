// 70. Climbing Stairs
// https://leetcode.com/problems/climbing-stairs/
//
// DP Top-down solution
package main

func climbStairs(n int) int {
	memo := make(map[int]int)

	var dp func(n int) int
	dp = func(n int) int {
		if n < 0 {
			return 0
		}
		if n == 0 {
			return 1
		}
		
		if _, ok := memo[n]; ok {
			return memo[n]
		}

		res := dp(n-1) + dp(n-2)
		memo[n] = res

		return res
	}

	return dp(n)
}
