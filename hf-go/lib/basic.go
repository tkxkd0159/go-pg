package lib

import "fmt"

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
