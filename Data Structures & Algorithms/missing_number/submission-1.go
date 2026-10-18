package missingnumber

func missingNumber(nums []int) int {
	distributed := make([]int, len(nums)+1)

	for _, val := range nums {
		distributed[val]++
	}

	for i := 0; i < len(distributed); i++ {
		if distributed[i] == 0 {
			return i
		}
	}

	return 0
}
