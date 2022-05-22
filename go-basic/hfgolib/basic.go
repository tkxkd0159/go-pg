package hfgolib

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

type Alltype struct {
	Num int64
	Dot float64
	Tf  bool
	Str string
	Arr []string
}

func (a *Alltype) AllAdd() []string {
	first := fmt.Sprintf("%f", float64(a.Num)+a.Dot)
	second := fmt.Sprintf("%t", a.Tf)
	third := first + second + a.Str
	fourth := append(a.Arr, third)

	return fourth
}

func GetAllTypes() Alltype {
	at := Alltype{}
	at.Num = 30
	at.Dot = 32.77
	at.Tf = true
	at.Str = "all type strcut"
	at.Arr = append(at.Arr, "I am")
	at.Arr = append(at.Arr, "returning")
	at.Arr = append(at.Arr, "all type struct")

	return at
}

func PrintSlice(s []int) string {
	return fmt.Sprintf("len=%d cap=%d %v\n", len(s), cap(s), s)
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
