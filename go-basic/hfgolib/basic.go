package hfgolib

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

func Hello() string {
	return "Let's learn Golang wiht Head First Go"
}

type Alltype struct {
	num int64
	dot float64
	tf  bool
	str string
	arr []string
}

func GetAllTypes() Alltype {
	at := Alltype{}
	at.num = 30
	at.dot = 32.77
	at.tf = true
	at.str = "all type strcut"
	at.arr = append(at.arr, "I am")
	at.arr = append(at.arr, "returning")
	at.arr = append(at.arr, "all type struct")

	return at
}

func PrintSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func ReturnMulti() (a int, b float64, c string, d []string) {
	a = 7
	b = 7.7
	c = "seven"
	d = []string{"string", "array"}
	return
}

func ReturnMulti2() (int, float64, string, []int) {
	return 7, 7.7, "seven", []int{12, 78, 50}
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type dirFlag bool

func PrintSystemInfo(flag dirFlag) {
	myOS, myArch, gover := runtime.GOOS, runtime.GOARCH, runtime.Version()
	fmt.Printf("%s / %s / %s ! \n", myOS, myArch, gover)

	if flag == true {
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("CWD is %s \n", path)

		files, err := ioutil.ReadDir(".")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("\n", "File list in ", path)
		for _, file := range files {
			fmt.Println(file.Name(), file.IsDir())
		}
	}

}
