// 209. Minimum Size Subarray Sum
// https://leetcode.com/problems/minimum-size-subarray-sum
//
// Sliding window (variable size), all nums positive. Expand r, then while the window is valid (sum >= target)
// record its length and shrink from the left. Record inside the while because validity is what we are minimizing over.
// O(n) time (each index enters and leaves once), O(1) space.
func minSubArrayLen(target int, nums []int) int {
    minLen, sum, l:= math.MaxInt32, 0, 0

    for r := 0; r < len(nums); r++ {
        sum += nums[r]
        for sum >= target {
            minLen = min(minLen, r - l + 1)
            sum -= nums[l]
            l++
        }
    }

    if minLen == math.MaxInt32 {
        return 0
    }

    return minLen
}
