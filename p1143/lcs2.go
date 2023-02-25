package p1143

func lcs2(text1 string, text2 string) int {
	var f = lcs2Finder{
		cache1: make([]int, len(text1)),
		cache2: make([]int, len(text2)),
		text1:  text1,
		text2:  text2,
	}
	for i := range f.cache1 {
		f.cache1[i] = -1
	}
	for i := range f.cache2 {
		f.cache2[i] = -1
	}
	return f.do(0, 0)
}

type lcs2Finder struct {
	cache1 []int
	cache2 []int
	text1  string
	text2  string
}

func (f *lcs2Finder) do(i, j int) int {
	if i >= len(f.text1) || j >= len(f.text2) {
		return 0
	}

	if f.text1[i] == f.text2[j] {
		return 1 + f.do(i+1, j+1)
	}

	lcSec1 := f.cache1[i+1]
	if lcSec1 == -1 {
		lcSec1 = f.do(i+1, j)
		f.cache1[i+1] = lcSec1
	}
	lcSec2 := f.cache1[j+1]
	if lcSec2 == -1 {
		lcSec2 = f.do(i, j+1)
		f.cache2[j+1] = lcSec1
	}

}
