// 239. Sliding Window Maximum
// https://leetcode.com/problems/sliding-window-maximum/

// O((n-k) * k) time complexity

package main

func maxSlidingWindow(nums []int, k int) []int {
	var maxes []int

	max := nums[0]
	maxIdx := 0
	for i, num := range nums {
		if maxIdx < i-k+1 {
			j := i - k + 1
			max = nums[j]
			for j <= i {
				if max <= nums[j] {
					max = nums[j]
					maxIdx = j
				}
				j++
			}
		}

		if max <= num {
			max = num
			maxIdx = i
		}

		if i < k-1 {
			continue
		}
		maxes = append(maxes, max)
	}

	return maxes
}
