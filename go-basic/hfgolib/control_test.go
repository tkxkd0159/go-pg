package hfgolib_test

import (
	"github.com/tkxkd0159/go-pg/go-basic/hfgolib"
	"testing"
)

var (
	CommonNum = 30
	CommonStr = "Lucky"
)

func TestBasicLoop(t *testing.T) {
	i := 1
	for i <= 3 {
		i = i + 1
	}
	if i != 4 {
		t.Errorf("First loop is wrong")
	}

	for j := 7; j <= 9; j++ {
		if j == 9 {
			t.Logf("Second loop is right, j == %d", j)
		}
	}

	k := 0
	for {
		k += 1
		if k == 4 {
			break
		}
		if k == 4 {
			t.Errorf("Third loop is wrong")
		}
	}

	chkRemain := [3]int{0, 2, 4}
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		for _, elem := range chkRemain {
			if n == elem {
				t.Errorf("Fourth loop is wrong")
			}
		}
	}
}

func TestForUnicode(t *testing.T) {
	var _ = CommonNum
	var _ = CommonStr

	tests := hfgolib.ForUnicode("abc가나다")
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
		hfgolib.ForUnicode("abc가나다")
	}
}
