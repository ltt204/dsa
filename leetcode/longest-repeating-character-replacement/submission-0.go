package main

// 424. Longest Repeating Character Replacement
// https://leetcode.com/problems/longest-repeating-character-replacement/description/
//
// number of replacements == length of substring - max frequency of a character in the substring
// To find the max, we can use an array of size 26 to store the frequency of each character in the substring. Then continue to comparing the max frequency
// The key here is when we shrink the window (l--), we decrease the frequency of the character at l and remain other.
// Why this work? it like we are counting frequency for the whole string, just it just when we shrink, we removed an element from the string.

func characterReplacement(s string, k int) int {
	counts := [26]int{}
	maxCount, l, best := 0, 0, 0

	for r := 0; r < len(s); r++ {
		counts[s[r]-'A']++
		maxCount = max(maxCount, counts[s[r]-'A'])
		if (r-l+1)-maxCount > k {
			counts[s[l]-'A']--
			l++
		}

		best = max(best, (r - l + 1))
	}

	return best
}
