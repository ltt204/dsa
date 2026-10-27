// 643. Maximum Average Subarray I
// https://leetcode.com/problems/maximum-average-subarray-i
//
// Sliding window (fixed size k): window [l, r] changes by one element per step,
// so add nums[r] and drop nums[l] once the window exceeds k. O(n) time, O(1) space.
func findMaxAverage(nums []int, k int) float64 {
    maxAvg := -math.MaxFloat32
    avg := 0.0

    l, r := 0, 0
    for r < len(nums) {
        avg += float64(nums[r])

        if r - l + 1 > k {
            avg -= float64(nums[l])
            l++
        }

        if r - l + 1 == k {
            maxAvg = max(maxAvg, avg/float64(k))
        }

        r++
    }

    return maxAvg
}
