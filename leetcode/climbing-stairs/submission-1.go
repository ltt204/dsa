// 70. Climbing Stairs
// https://leetcode.com/problems/climbing-stairs/
//
// DP bottom-up solution
package main

func climbStairs(n int) int {
	memo := make(map[int]int)
	memo[0] = 1
	memo[1] = 1

	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}

	return memo[n]
}
