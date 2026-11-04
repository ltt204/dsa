package longestpalindromesubstring

func longestPalindrome(s string) string {
	memo := make([]string, len(s)+1)

	isPali := func(s string) bool {
		if s == "" {
			return false
		}
		l, r := 0, len(s)-1
		for l < r {
			if s[l] != s[r] {
				return false
			}

			l++
			r--
		}

		return true
	}

	var dp func(s string) string
	dp = func(s string) string {
		if s == "" {
			return s
		}

		if memo[len(s)] != "" {
			return memo[len(s)]
		}

		if isPali(s) {
			memo[len(s)] = s
			return s
		}

		resLeft := dp(s[:len(s)-1])
		resRight := dp(s[1:])

		if resLeft != "" {
			memo[len(resLeft)] = resLeft
		}

		if resRight != "" {
			memo[len(resRight)] = resRight
		}

		return ""
	}

	dp(s)

	for i := len(s); i >= 0; i-- {
		if memo[i] != "" {
			return memo[i]
		}
	}

	return memo[1]
}
