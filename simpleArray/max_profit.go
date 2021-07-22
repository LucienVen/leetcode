// @Author : a1234
// @Time : 2021/7/22 5:33 下午 
// @Describe :  股票最大利润

package simpleArray

import "math"

func MaxProfit(prices []int) int {
	if len(prices) < 2 || prices == nil {
		return 0
	}

	length := len(prices)

	// 初始条件
	var (
		hold = -prices[0] // 持有股票
		noHold = 0 // 没持有股票
	)

	for i := 1; i < length; i++ {
		noHold = int(math.Max(float64(noHold), float64(hold+prices[i])))
		hold = int(math.Max(float64(hold), float64(noHold-prices[i])))
	}

	return noHold
}

// 贪心算法
func GreedMaxProfit(prices []int) int {
	if len(prices) < 2 || prices == nil {
		return 0
	}

	var total = 0
	for i := 0; i < len(prices)-1; i++ {
		
		if (prices[i+1] - prices[i]) > 0 {
			total = total +  (prices[i+1] - prices[i])
		}
	}

	return total
}