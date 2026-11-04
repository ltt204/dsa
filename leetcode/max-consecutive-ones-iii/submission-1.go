// https://leetcode.com/problems/max-consecutive-ones-iii/
//
// Core idea: Sliding window
// Counting 0s by continuouly decrease k, if k < 0, move left pointer to right until k >= 0

package maxconsecutiveonesiii

func longestOnes(nums []int, k int) int {
	l := 0

	for _, num := range nums {
		if num == 0 {
			k--
		}

		if k < 0 {
			if nums[l] == 0 {
				k++
			}
			l++
		}
	}

	return len(nums) - l
}
