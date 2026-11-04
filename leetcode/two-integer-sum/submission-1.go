func twoSum(nums []int, target int) []int {
    var m = make(map[int]int, len(nums))

	for idx, val := range nums {
		if _, ok := m[target-val]; ok {
			return []int{m[target-val], idx}
		}

		m[val] = idx
	}

	return nil
}
