// 17. Letter Combinations of a Phone Number
// https://leetcode.com/problems/letter-combinations-of-a-phone-number
package main

import "bytes"

func letterCombinations(digits string) []string {
	numToLet := map[int]string{
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}

	var result []string
	var path bytes.Buffer

	var backtrack func(digIdx int)
	backtrack = func(digIdx int) {
		if digIdx >= len(digits) {
			result = append(result, path.String())
			return
		}

		lets := numToLet[int(digits[digIdx]-'0')]
		for i := 0; i < len(lets); i++ {
			path.WriteByte(lets[i])
			backtrack(digIdx + 1)
			path.Truncate(path.Len() - 1)
		}
	}

	backtrack(0)

	return result
}
