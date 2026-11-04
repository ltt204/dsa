func isPalindrome(s string) bool {
	// check if empty or 1 char only
	// 1. lower case + remove spaces	
	// Two poitners:
	// one go forward and one go backward
	// mismatch => not palindrome

	s = strings.ToLower(s)
	if len(s) == 0 || len(s) == 1 {
		return true
	}
	left := 0
	right := len(s) - 1
	
	for left < right {
		if !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])) {
			left++
			continue
		}
		if !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])) {
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}