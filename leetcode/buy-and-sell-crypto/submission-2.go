func maxProfit(prices []int) int {
    left, right := 0 , 1
    maxProfit := 0

    for right < len(prices) {
        if prices[left] > prices[right] {
            left = right
            right ++
            continue
        }
        
        if maxProfit < (prices[right] - prices[left]) {
            maxProfit = prices[right] - prices[left]
        } 
        
        right ++
    }

    return maxProfit
}