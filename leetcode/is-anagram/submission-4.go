func isAnagram(s string, t string) bool {
	if (len(s) != len(t)) {return false}
	
	var m = make(map[rune]int)

	for _, char := range s {
		m[char]++
	}

	for _, char := range t {
		m[char]--
	}

	for _, val := range m {
		if val != 0 {
			return false
		}
	}

	return true
}