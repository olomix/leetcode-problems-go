package p72

type cacheKey struct {
	word1 string
	word2 string
}

var cache = make(map[cacheKey]int)

type cacheKey2 [2]int16

func minDistance(word1 string, word2 string) int {
	if word1 == word2 {
		return 0
	}

	if word1 == "" {
		return len(word2)
	}

	if word2 == "" {
		return len(word1)
	}

	if word1[0] == word2[0] {
		return minDistance(word1[1:], word2[1:])
	}

	ck := cacheKey{
		word1: word1,
		word2: word2,
	}
	if v, ok := cache[ck]; ok {
		return v
	}

	case1N := minDistance(word1[1:], word2)
	case2N := minDistance(word2[:1]+word1[1:], word2)
	case3N := minDistance(word2[:1]+word1, word2)

	if case2N < case1N {
		case1N = case2N
	}
	if case3N < case1N {
		case1N = case3N
	}
	cache[ck] = case1N + 1
	return case1N + 1
}

func minDistance2(word1 string, word2 string) int {
	return minDistanceRecur(word1, word2, len(word1), len(word2), nil)
}
func minDistanceRecur(word1 string, word2 string, word1Ln, word2Ln int,
	cache map[cacheKey2]int) int {

	if word1Ln == 0 {
		return word2Ln
	}

	if word2Ln == 0 {
		return word1Ln
	}

	if word1[word1Ln-1] == word2[word2Ln-1] {
		return minDistanceRecur(word1, word2, word1Ln-1, word2Ln-1, cache)
	}

	ck := cacheKey2{int16(word1Ln), int16(word2Ln)}
	if cache == nil {
		cache = make(map[cacheKey2]int)
	}
	if v, ok := cache[ck]; ok {
		return v
	}

	caseReplace := minDistanceRecur(word1, word2, word1Ln-1, word2Ln-1, cache)
	caseDelete := minDistanceRecur(word1, word2, word1Ln-1, word2Ln, cache)
	caseInsert := minDistanceRecur(word1, word2, word1Ln, word2Ln-1, cache)

	min := caseDelete
	if caseInsert < min {
		min = caseInsert
	}
	if caseReplace < min {
		min = caseReplace
	}
	cache[ck] = min + 1
	return min + 1
}
