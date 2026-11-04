// 198. House Robber
// https://leetcode.com/problems/house-robber
//
// At house 0, we have 1 option is robbed
// At house 1, we have only 1 option which is rob house 1 + unrob house 0 - or keep house 0
// But from house 2, we can see we have two option:
//     1. Not rob the house 0, and robbed the current house
//     2. Skip current hourse
// If we take the current house, we can add the value the nearest not adjacent hourse.
// Since we already calculated from the very first, the add will also contains those very first house
// - accumulate and using max will help us find the max at the current house

package robbinghouse

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	memo := make(map[int]int)
	memo[0] = nums[0]
	memo[1] = max(memo[0], nums[1])

	for i := 2; i < len(nums); i++ {
		memo[i] = max(memo[i-1], nums[i]+memo[i-2])
	}

	return memo[len(memo)-1]
}
