func lengthOfLongestSubstring(s string) int {
    if len(s) < 2 {
        return len(s)
    }

    hmap := make(map[rune]int)
    l := 0
    maxLen := 0

    for r, val := range s {
        if _, ok := hmap[val]; ok {
            if hmap[val] >= l {
                l = hmap[val] + 1
            }
        }
     
        hmap[val] = r

        if maxLen < (r - l + 1) {
            maxLen = (r - l + 1)
        }

    }

    return maxLen
}
