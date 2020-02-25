package bestTimeToBuyAndSellStock

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	wants := []int{
		5,
		0,
	}
	in := [][]int{
		{7, 1, 5, 3, 6, 4},
		{7, 6, 4, 3, 1},
	}
	for k, v := range wants {
		result := MaxProfit(in[k])
		if result != v {
			t.Errorf("in: %v\twant: %v\tresult: %v\n", in[k], v, result)
		}
	}
}
