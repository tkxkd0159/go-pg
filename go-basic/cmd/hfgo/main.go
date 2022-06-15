package main

import (
	"fmt"
	hg "github.com/tkxkd0159/go-pg/go-basic/hfgolib"
)

func main() {
	tt := hg.StructSample{}
	tt.SetE1Wptr("aaaa")
	tt.SetE2Wptr(77)
	fmt.Println(tt)
	tt = tt.SetE1("bbbb")
	tt = tt.SetE2(22)
	fmt.Println(tt)
}
