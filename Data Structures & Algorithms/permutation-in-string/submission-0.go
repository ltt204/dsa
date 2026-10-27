package permutationinstring

// 567. Permutation in String
// https://leetcode.com/problems/permutation-in-string/description/
//
// Basically, s2 contains one of s1's permutation.
// Using the idea of longest-repeating-character-replacement, Im thinking about:
// 1. Intializa the counts array with chartacters frequency of s1
// 2. Iterate through s2, with l and r. When the (r-l+1) > len(s1), we shrink the window by l++ and decrease the frequency of s2[l] in counts array.
//    Along with this, using a sum variable, which tracking the total value of counts array, if sum == 0, we found a permutation of s1 in s2.
// 3. Iterating: We can skip the characters if counts[s[r]-'a'] == 0. Because it doesnt appear in s1 anyway
// O(m*n) solution

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	// m*n

	permutation := func(l, sum int, counts [26]int) bool {
		for r := l; r < len(s2); r++ {
			if counts[s2[r]-'a'] == 0 {
				break
			}

			counts[s2[r]-'a']--
			sum--
		}

		if sum == 0 {
			return true
		}

		return false
	}

	counts := [26]int{}
	sum := 0
	for i := 0; i < len(s1); i++ {
		counts[s1[i]-'a']++
		sum++
	}

	for l := 0; l < len(s2); l++ {
		if counts[s2[l]-'a'] == 0 {
			continue
		}

		if permutation(l, sum, counts) {
			return true
		}
	}

	return false
}
