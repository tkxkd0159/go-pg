package hfgolib_test

import (
	"github.com/tkxkd0159/go-pg/go-basic/hfgolib"
	"testing"
)

func TestPrintSlice(t *testing.T) {
	got := hfgolib.PrintSlice([]int{1, 2, 3, 4, 5})
	exp := "len=5 cap=5 [1 2 3 4 5]\n"
	if got != exp {
		t.Error("TestPrintSlice is failed")
	}
}
