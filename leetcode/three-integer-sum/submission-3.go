func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var result = make(map[[3]int]struct{})

	for idx, val := range nums {
		// only 2 elements left
		if idx == len(nums) - 2 {
			break
		}

		left := idx + 1
		right := len(nums)-1
		for left < right {
			sum := val + nums[left] + nums[right]

			if sum == 0 {
				result[[3]int{val, nums[left], nums[right]}] = struct{}{}
				left++
				right--
				continue	
			}

			if sum < 0 {
				left++
				continue
			}

			if sum > 0 {
				right--
				continue
			}
		}
	} 

	var result2D [][]int
	for key,_ := range result {
		result2D = append(result2D, key[:])
	}

	return result2D
}
