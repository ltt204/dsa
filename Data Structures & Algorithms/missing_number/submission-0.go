package missingnumber

import "slices"

func missingNumber(nums []int) int {
	slices.Sort(nums)

	for idx, val := range nums {
		if idx != val {
			return idx
		}
	}

	return len(nums)
}
