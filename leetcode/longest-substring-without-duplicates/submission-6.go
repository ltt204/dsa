func lengthOfLongestSubstring(s string) int {

    // if len(s) < 2 {
    //     return len(s)
    // }

    var hmap = make(map[byte]int)
    maxLen := 0
    i := 0

    for i < len(s) {
        b := s[i]
        if _, ok := hmap[b]; ok {
            if maxLen < len(hmap) {
                maxLen = len(hmap)
            }
            i = hmap[b] + 1
            hmap = make(map[byte]int)
        }
        hmap[s[i]] = i
        i ++
    }

    if maxLen < len(hmap) {
        maxLen = len(hmap)
    }


    return maxLen
}
