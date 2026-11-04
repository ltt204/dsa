package permutationinstring

// 567. Permutation in String
// https://leetcode.com/problems/permutation-in-string/description/
//
// O(m+n) solution
// Still counting, but when the window hit the size of s1, we check if the counts array is valid.
// Iterating through 26 elements is O(1) time complexity, so the overall time complexity is O(m+n)

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	counts := [26]int{}
	for i := 0; i < len(s1); i++ {
		counts[s1[i]-'a']++
	}

	l := 0
	counts2 := [26]int{}
	for r := 0; r < len(s2); r++ {
		counts2[s2[r]-'a']++

		if r-l+1 < len(s1) {
			continue
		}

		if r-l+1 > len(s1) {
			counts2[s2[l]-'a']--
			l++
		}

		if r-l+1 == len(s1) {
			valid := true
			for i := 0; i < len(counts); i++ {
				if counts[i] != 0 && counts[i] != counts2[i] {
					valid = false
					break
				}
			}

			if valid {
				return true
			}
		}
	}

	return false
}
