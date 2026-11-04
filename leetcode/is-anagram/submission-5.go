func isAnagram(s string, t string) bool {
	if (len(s) != len(t)) {return false}
	
	// var m = make(map[rune]int)
	var counts [26]int

	for _, char := range s {
		counts[char-'a']++
	}

	for _, char := range t {
		counts[char-'a']--
	}

	for _, val := range counts {
		if val != 0 {
			return false
		}
	}

	return true
}