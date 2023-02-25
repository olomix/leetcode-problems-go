package p1143

func longestCommonSubsequence(text1 string, text2 string) int {
	cache := make(map[[2]string]int)
	return lcs(text1, text2, cache)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lcs(text1, text2 string, cache map[[2]string]int) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	if text1[0] == text2[0] {
		return 1 + lcs(text1[1:], text2[1:], cache)
	}

	key1 := [2]string{text1[1:], text2}
	lcs1, ok := cache[key1]
	if !ok {
		lcs1 = lcs(text1[1:], text2, cache)
		cache[key1] = lcs1
	}
	key2 := [2]string{text1, text2[1:]}
	lcs2, ok := cache[key2]
	if !ok {
		lcs2 = lcs(text1, text2[1:], cache)
		cache[key2] = lcs2
	}
	return max(lcs1, lcs2)
}
