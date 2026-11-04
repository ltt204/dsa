// https://leetcode.com/problems/max-consecutive-ones-iii/
//
// Core idea: Sliding window
// Counting 1, lenWindow - count1 > k, move left pointer

package maxconsecutiveonesiii

func longestOnes(nums []int, k int) int {
	l, r := 0, 0
	maxLen, count1 := 0, 0
	for r < len(nums) {

		if nums[r] == 1 {
			count1++
		}

		for (r-l+1)-count1 > k {
			if count1 > 0 && nums[l] == 1 {
				count1--
			}

			l++
		}
		maxLen = max(maxLen, r-l+1)
		r++
	}

	return maxLen
}
