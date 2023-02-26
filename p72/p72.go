package p72

type cacheKey struct {
	word1 string
	word2 string
}

var cache = make(map[cacheKey]int)

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
