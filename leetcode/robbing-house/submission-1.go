// 198. House Robber
// https://leetcode.com/problems/house-robber
//
// Same idea with submission-0, which is bottom-up DP. With this one we geos from the top - dp(len(nums)-1)

package robbinghouse

func rob(nums []int) int {
	memo := make(map[int]int)

	var dp func(i int) int
	dp = func(i int) int {
		if i < 0 {
			return 0
		}
		if _, ok := memo[i]; ok {
			return memo[i]
		}

		res := max(dp(i-1), nums[i]+dp(i-2))
		memo[i] = res

		return res
	}

	dp(len(nums) - 1)

	return memo[len(nums)-1]
}
