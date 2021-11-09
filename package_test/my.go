package main

import (
	"fmt"

	"github.com/tkxkd0159/go-server/package_test/pkg1"
	sub2 "github.com/tkxkd0159/go-server/package_test/pkg1/pkg2"
	newsub1 "github.com/tkxkd0159/go-server/package_test/pkg3"

	"golang.org/x/example/stringutil"
)

func main() {
	fmt.Println(" * Package Import Test")
	fmt.Println(pkg1.Hello())
	fmt.Println(pkg1.Hello11())
	fmt.Println(sub2.Hello())
	fmt.Println(newsub1.Hello())
	fmt.Println(stringutil.Reverse("!selpmaxe oG ,olleH"))
}