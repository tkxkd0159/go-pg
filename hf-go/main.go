package main

import (
	"fmt"
	jslib "hf-go/lib"
	"log"
	"os"
	"runtime"
)

func main() {
	fmt.Println(jslib.Hello())
	fmt.Printf("%+v\n", jslib.GetAllTypes())
	fmt.Println(jslib.ReturnMulti())
	fmt.Println(jslib.ReturnMulti2())

	myOS, myArch, gover := runtime.GOOS, runtime.GOARCH, runtime.Version()
	fmt.Printf("%s / %s / %s ! \n", myOS, myArch, gover)

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
}
