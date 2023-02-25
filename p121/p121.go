package p121

// that was correct but too slow
func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	profits := make([]int, len(prices)-1)
	for i := 0; i < len(profits); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > profits[i] {
				profits[i] = prices[j] - prices[i]
			}
		}
	}

	m := profits[0]
	for i := 1; i < len(profits); i++ {
		if profits[i] > m {
			m = profits[i]
		}
	}

	return m
}

func maxProfit(prices []int) int {
	min1 := prices[0]
	max1 := min1
	newMin := -1

	for i := 1; i < len(prices); i++ {
		if prices[i] > max1 {
			max1 = prices[i]
		}

		if newMin != -1 && prices[i]-newMin > max1-min1 {
			min1 = newMin
			max1 = prices[i]
			newMin = -1
		}
		if prices[i] < min1 && (newMin == -1 || prices[i] < newMin) {
			newMin = prices[i]
		}
	}

	return max1 - min1
}
