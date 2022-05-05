package main

import (
	"fmt"
	jslib "github.com/tkxkd0159/go-pg/go-basic/hfgolib"
)

func main() {
	fmt.Println(jslib.Hello())
	fmt.Printf("%+v\n", jslib.GetAllTypes())
	fmt.Println(jslib.ReturnMulti())
	fmt.Println(jslib.ReturnMulti2())
	fmt.Printf("\n")

	jslib.PrintSystemInfo()

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	k := 0
	for {
		fmt.Println("loop")
		k += 1
		if k == 4 {
			break
		}
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
