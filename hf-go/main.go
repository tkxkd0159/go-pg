package main

import (
	"fmt"
	jslib "hfgo/lib"
)

func main() {
	fmt.Println(jslib.Hello())
	fmt.Printf("%+v\n", jslib.Types())
	fmt.Println(jslib.ReturnMulti())
	fmt.Println(jslib.ReturnMulti2())
}
