func lengthOfLongestSubstring(s string) int {
    if len(s) < 2 {
        return len(s)
    }

    var lastSeen [256]int
	for i := range lastSeen {
		lastSeen[i] = -1
	}

    l := 0
    maxLen := 0

    for r, val := range s {
        if lastSeen[val] >= l {
            l = lastSeen[val] + 1
        }
     
        lastSeen[val] = r

        if maxLen < (r - l + 1) {
            maxLen = (r - l + 1)
        }

    }

    return maxLen
}
