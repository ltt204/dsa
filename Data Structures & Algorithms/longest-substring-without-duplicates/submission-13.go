// 3. Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters
//
// Sliding window (variable size) with lastSeen[c] = index+1 of c's last occurrence.
// Invariant: s[l..r] has no duplicate. Storing r+1 makes the jump `l = max(l, lastSeen)`
// equal to `max(l, j+1)`; the max keeps l from moving backward when the old copy is outside the window.
func lengthOfLongestSubstring(s string) int {
    lastSeen := make(map[byte]int)

    l, r := 0, 0
    maxLen := 0
    for r < len(s) {
        if _, ok := lastSeen[s[r]]; ok {
            l = max(l, lastSeen[s[r]])
        }

        maxLen = max(maxLen, r-l+1)
        lastSeen[s[r]] = r+1

        r++
    }

    return maxLen
}
