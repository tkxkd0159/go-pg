package hfgolib

import (
	"testing"
)

var (
	CommonNum = 30
	CommonStr = "Lucky"
)

func TestForUnicode(t *testing.T) {
	var _ = CommonNum
	var _ = CommonStr

	tests := ForUnicode("abc가나다")
	ans := []struct {
		codePoint rune
		char      string
	}{
		{97, "a"},
		{98, "b"},
		{99, "c"},
		{44032, "가"},
		{45208, "나"},
		{45796, "다"},
	}

	for _, a := range ans {
		if tests[a.char] != a.codePoint {
			t.Errorf("%d is not equal to this unicode point %s", tests[a.char], a.char)
		}
	}
}

func BenchmarkForUnicode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ForUnicode("abc가나다")
	}
}
