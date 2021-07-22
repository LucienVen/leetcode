// @Author : a1234
// @Time : 2021/7/22 5:42 下午 
// @Describe :  

package simpleArray

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7,1,5,3,6,4}
	maxProfit := MaxProfit(prices)
	fmt.Println("max_profit:", maxProfit)
}

func TestGreedMaxProfit(t *testing.T) {
	prices := []int{7,1,5,3,6,4}
	maxProfit := GreedMaxProfit(prices)
	fmt.Println("max_profit:", maxProfit)
}