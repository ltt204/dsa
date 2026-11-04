func twoSum(nums []int, target int) []int {
    var m = make(map[int]int)
	var result = []int{-1,-1}

	for idx, val := range nums {
		if _, ok := m[target-val]; ok {
			result[0]=m[target-val]
			result[1]=idx

			return result
		}

		m[val] = idx
	}

	return result
}
