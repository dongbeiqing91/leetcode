package easy

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/

func maxProfit(prices []int) int {
	lowestPrice := prices[0]
	profit := 0
	mProfit := 0
	for _, p := range prices {
		if p < lowestPrice {
			lowestPrice = p
		} else {
			profit = p - lowestPrice
			if profit > mProfit {
				mProfit = profit
			}
		}
	}
	return mProfit
}
