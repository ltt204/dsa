// Sorting and Searching - Concert Tickets
//
// The idea for tracking unsold is the right one, but the implementation is a bit off.
// We can see that the path for "finding the valid unsold ticket" is taking, at worst case,
// O(n) time complexity.
// The idea for improvment is using prefix sum. Here is the plan:
// The unsold array is now tracked as a int array, where unsold[i] is default with 1.
// We can tell that "The number of unsold tickets is the total of unsold tickets from 0 to i"
// is the prefix sum of unsold array.
// For example: prices: [3, 5, 7, 9], unsold: [1, 0, 1, 1], prefix sum: [1, 2, 3, 4]
// For getting the number of unsold tickets for 7, we can see that the index of 7 is 2,
// and the prefix sum of unsold[2] is 3, which means there are 3 unsold tickets from 0 to 2.
package main

import (
	"fmt"
	"slices"
	"sort"
)

func solve(n int, m int, prices []int, members []int) []int {
	var result []int
	sold := make([]bool, len(prices))

	slices.Sort(prices)
	for i := range m {
		index := sort.Search(n, func(idx int) bool {
			return prices[idx] > members[i]
		})

		for index >= 0 {
			if index >= n {
				index--
				continue
			}

			if !sold[index] && prices[index] <= members[i] {
				result = append(result, prices[index])
				sold[index] = true
				break
			}

			index--
		}

		if index == -1 {
			result = append(result, -1)
		}
	}

	return result
}

func main() {
	// Read input
	var n, m int
	fmt.Scan(&n, &m)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}

	result := solve(n, m, a, b)

	// Print output
	for _, price := range result {
		fmt.Println(price)
	}
}
