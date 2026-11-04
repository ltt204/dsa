func groupAnagrams(strs []string) [][]string {
	var m = make(map[[26]int][]string)

	for _, str := range strs {
		var counts [26]int

		for _, char := range str {
			counts[char - 'a']++
		}

		m[counts] = append(m[counts], str)
	}

	var result = make([][]string, 0, len(m))
	for _, val := range m {
		row := make([]string, 0, len(val)+1)
		row = append(row, val...)

		result = append(result, row)
	} 

	return result
}
