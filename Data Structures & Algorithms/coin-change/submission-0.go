// 322. Coin Change
// https://leetcode.com/problems/coin-change
//
// Focus on the amount. At a specific state of the amount --> have a min coins to make exactly that amount.
// Example: 1, 2, 5 -- amount=8
// Suppose we have amount:
//   0 - 0 coins -- base case
//   1 - 1 coins -- can only take 1
//   2 - 1 coins -- we can either take 2 coins 1 or 1 coin 2, since the reqiurement is fewest --> 1 coin 2
//   3 - 2 coins -- either 3 coins 1 or 1 coin 1 and 1 coin 2'
//   .... And so on
// We can see that, at the coin 3, it is the combination of the current coin (only 1 coin) we are processing (either 1 or 2) with the min coins of the left over amount,
// suppose we are working with coin 1, we will have fomula like this: 1 (current coin 1) + memo[3 - 1] => 1 + memo[2] => 1 + 1 = 2 coins
// similar with coin 2, we will have: 1 + memo[3-2] => 1 + memo[1] = 2 coins
// finally we will have this fomula:
// memo[current_amount] = min(memo[current_amount], 1 + memo[leftover = current_amount - coins[current_index]])

package coinchange

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}
