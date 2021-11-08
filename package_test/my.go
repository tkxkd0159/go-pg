package main

import (
	"fmt"

	sub1 "github.com/tkxkd0159/package_test/pkg1"
	sub2 "github.com/tkxkd0159/package_test/pkg1/pkg2"
	newsub1 "github.com/tkxkd0159/package_test/pkg3"
)

func main() {
	fmt.Println(" * Package Import Test")
	fmt.Println(sub1.Hello())
	fmt.Println(sub1.Hello11())
	fmt.Println(sub2.Hello())
	fmt.Println(newsub1.Hello())
}