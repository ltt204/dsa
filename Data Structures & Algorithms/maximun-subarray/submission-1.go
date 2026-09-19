// 53. Maximun Subarray
// https://leetcode.com/problems/maximum-subarray
//
// The fomala: max(nums[i], nums[i] + dp[i-1]) and the base is dp[0]=nums[0]
// The idea: subarray is a contiguous non-empty sequence. Therefor, we can iterate the whole array and calculate for the maximum
// Example: [-2,1,-3,4,-1]
// 			dp[0] = -2
// 			dp[1] = max(1,-1) 	= 1
// 			dp[2] = max(-3,-2) 	= -2
// 			dp[3] = max(4,2) 	= 4 ---> From index 0 to 3, we accumulated all values and have -2, then we compare either -2 + 4 and the 4 itself greater.
// 										 Clearly we see the 4 is greate, taking it analogous to we drop the previous sequence (index 0 - 3) --> reset the sequence
// 			dp[4] = max(-1,3)	= 3
// Note 1: The previous need to sort, it turn the time complexity from O(n) to O(nlogn). Use a variable to hold the actual result instead of sorting

package maximunsubarray

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], nums[i]+dp[i-1])
		result = max(dp[i], result)
	}

	return result
}
