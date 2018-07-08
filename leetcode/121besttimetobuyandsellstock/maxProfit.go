package main

func maxProfit(prices []int) int {
	l := len(prices)
	if prices == nil || l < 2 {
		return 0
	}

	i, ret, curLow := 0, 0, prices[0]

	for i < l {
		// fmt.Printf("%d curLow %d %v\n", i, curLow, prices)
		if prices[i] < curLow {
			curLow = prices[i]
			continue
		}

		if (prices[i] - curLow) > ret {
			ret = (prices[i] - curLow)
		}
		i++
	}
	return ret

}
