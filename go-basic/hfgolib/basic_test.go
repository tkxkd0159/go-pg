package hfgolib_test

import (
	"github.com/tkxkd0159/go-pg/go-basic/hfgolib"
	"testing"
)

func TestIntMinBasic(t *testing.T) {
	ans := hfgolib.IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
	ans2 := hfgolib.IntMin(-2, 2)
	if ans2 != -2 {
		t.Errorf("IntMin(-2, 2) = %d; want -2", ans2)
	}
}

func TestPrintSlice(t *testing.T) {
	got := hfgolib.PrintSlice([]int{1, 2, 3, 4, 5})
	exp := "len=5 cap=5 [1 2 3 4 5]\n"
	if got != exp {
		t.Error("TestPrintSlice is failed")
	}
}
