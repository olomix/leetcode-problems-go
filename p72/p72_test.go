package p72

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minDistance(t *testing.T) {
	require.Equal(t, 3, minDistance("horse", "ros"))
	require.Equal(t, 5, minDistance("intention", "execution"))
	require.Equal(t, 6,
		minDistance("dinitrophenylhydrazine", "acetylphenylhydrazine"))
}

// 145185
// 144892
func Benchmark_minDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cache = make(map[cacheKey]int)
		require.Equal(b, 3, minDistance("horse", "ros"))
		require.Equal(b, 5, minDistance("intention", "execution"))
		require.Equal(b, 6,
			minDistance("dinitrophenylhydrazine", "acetylphenylhydrazine"))

	}
}

// 13254
func Benchmark_minDistance2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cache = make(map[cacheKey]int)
		require.Equal(b, 3, minDistance2("horse", "ros"))
		require.Equal(b, 5, minDistance2("intention", "execution"))
		require.Equal(b, 6,
			minDistance2("dinitrophenylhydrazine", "acetylphenylhydrazine"))
	}
}

// horse
// ros
// horse -> orse -> rorse -> rose -> ros
// dinitrop => acetyl

// orse | rorse | rhorse

// dinitro phenylhydrazine
// acetyl phenylhydrazine

// acetylphenylhydrazine -> dinitrophenylhydrazine
