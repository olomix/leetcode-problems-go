package p1143

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	testCases := []struct {
		text1, text2 string
		want         int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
		{"oxcpqrsvwf", "shmtulqrypy", 2},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			210},
		{"mhunuzqrkzsnidwbun", "szulspmhwpazoxijwbq", 6},
	}

	for _, tc := range testCases {
		got := longestCommonSubsequence(tc.text1, tc.text2)
		if got != tc.want {
			t.Errorf("longestCommonSubsequence(%v, %v) = %v, want %v", tc.text1,
				tc.text2,
				got, tc.want)
		}
	}
}
