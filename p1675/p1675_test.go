package p1675

import (
	"bufio"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestP1675(t *testing.T) {
	require.Equal(t, 1, minimumDeviation([]int{1, 2, 3, 4}))
	require.Equal(t, 3, minimumDeviation([]int{4, 1, 5, 20, 3}))
	require.Equal(t, 3, minimumDeviation([]int{2, 10, 8}))
	require.Equal(t, 5, minimumDeviation([]int{4, 9, 4, 5}))
	require.Equal(t, 0, minimumDeviation([]int{8, 1, 2, 1}))
	require.Equal(t, 5, minimumDeviation([]int{8, 4, 10, 4, 9}))
	require.Equal(t, 329, minimumDeviation([]int{136, 465, 87}))
	require.Equal(t, 1, minimumDeviation([]int{3, 5}))
	require.Equal(t, 2, minimumDeviation([]int{10, 4, 3}))
	require.Equal(t, 236, minimumDeviation([]int{610, 778, 846, 733, 395}))
	require.Equal(t, 99974096,
		minimumDeviation(readInts("testdata/99974096.txt")))
	require.Equal(t, 99932573,
		minimumDeviation(readInts("testdata/99932573.txt")))
	require.Equal(t, 0, minimumDeviation(readInts("testdata/0.txt")))
	require.Equal(t, 99977690,
		minimumDeviation(readInts("testdata/99977690.txt")))
}

func readInts(fName string) []int {
	f, err := os.Open(fName)
	if err != nil {
		panic(err)
	}
	var r []int
	s := bufio.NewScanner(f)
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}
		r = append(r, i)
	}
	return r
}
