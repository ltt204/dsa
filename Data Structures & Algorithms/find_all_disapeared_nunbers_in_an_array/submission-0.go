package findalldisapearednunbersinanarray

func findDisappearedNumbers(nums []int) []int {
	distributed := make([]int, len(nums)+1)

	for _, val := range nums {
		distributed[val]++
	}

	var result []int
	for i := 1; i < len(distributed); i++ {
		if distributed[i] == 0 {
			result = append(result, i)
		}
	}

	return result
}
