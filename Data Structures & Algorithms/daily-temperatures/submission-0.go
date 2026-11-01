// Source: https://leetcode.com/problems/daily-temperatures/
//
// Using double ended queue to store the index of the temperatures.
// If the current temperature is greater than the temperature at the index stored in the deque,
// we pop the index from the deque and calculate the difference between the current index and the popped index to get the number of days until a warmer temperature.
// We repeat this process until we find a temperature that is not warmer than the current temperature or until the deque is empty. Finally, we append the current index to the deque.
//
// Time complexity: O(n), where n is the length of the temperatures array. Each index is pushed and popped from the deque at most once.
package main

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	var deque []int

	for i, tem := range temperatures {
		for len(deque) > 0 && tem > temperatures[deque[len(deque)-1]] {
			result[deque[len(deque)-1]] = i - deque[len(deque)-1]
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)
	}

	return result
}
