package p402

import (
	"strings"
)

func removeKdigits(num string, k int) string {
	if k <= 0 {
		return strings.TrimLeft(num, "0")
	}

	numB := []byte(num)

	for k > 0 {
		if k >= len(numB) {
			return "0"
		}

		foundBest := false
		for i := 0; i < len(numB)-1; i++ {
			if numB[i] > numB[i+1] {
				foundBest = true
				numB = append(numB[:i], numB[i+1:]...)
				break
			}
		}
		if !foundBest {
			numB = numB[:len(numB)-1]
		}
		for len(numB) > 0 && numB[0] == '0' {
			numB = numB[1:]
		}
		k--
	}

	if len(numB) == 0 {
		return "0"
	}

	return string(numB)
}
