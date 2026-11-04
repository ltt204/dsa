// 131. Palindrome Partitioning
// http://leetcode.com/problems/palindrome-partitioning
//
// "Partition" part is homing on in distract us. This still a normal backtracking problem
// since every time we backtrack with r+1 and the left (l) is still kept, we will get the string[l:r+1] as our new partition
// So we can tacitly implicit understand that the partition will always be create, and we dont have to track it.
package palindromepartitioning

func partition(s string) [][]string {
	var result [][]string
	var path []string

	isPali := func(s string, l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}

			l++
			r--
		}

		return true
	}

	var backtrack func(l int)
	backtrack = func(l int) {
		if l >= len(s) {
			temp := append([]string{}, path...)
			result = append(result, temp)
			return
		}

		for r := l; r < len(s); r++ {
			if !isPali(s, l, r) {
				continue
			}

			path = append(path, s[l:r+1])
			backtrack(r + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)

	return result
}
