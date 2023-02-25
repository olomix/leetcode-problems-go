package p739

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDailyTemperatures(t *testing.T) {
	hugeInput := make([]int, 100000)
	hugeWant := make([]int, 100000)
	for i := range hugeInput {
		hugeInput[i] = 99
		hugeWant[len(hugeWant)-1-i] = i
	}
	hugeInput[99999] = 100

	testCases := []struct {
		in   []int
		want []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{[]int{99, 99, 99, 100}, []int{3, 2, 1, 0}},
		{[]int{34, 80, 80, 34, 34, 80, 80, 80, 80, 34},
			[]int{1, 0, 0, 2, 1, 0, 0, 0, 0, 0}},
		{hugeInput, hugeWant},
	}

	for _, tc := range testCases {
		got := dailyTemperatures(tc.in)
		require.Equal(t, tc.want, got)
	}
}
