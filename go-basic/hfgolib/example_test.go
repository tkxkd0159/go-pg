package hfgolib_test

import (
	"fmt"
	"github.com/tkxkd0159/go-pg/go-basic/hfgolib"
)

func ExampleAlltype_AllAdd() {
	mock := hfgolib.Alltype{Num: 2, Dot: 0.29, Tf: true, Str: "extarString", Arr: []string{"initial member"}}
	mock2 := hfgolib.Alltype{}
	fmt.Println(mock.AllAdd())
	fmt.Println(mock2.AllAdd())
	// Output:
	// [initial member 2.290000trueextarString]
	// [0.000000false]
}
