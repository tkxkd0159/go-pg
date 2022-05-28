package hfgolib

import (
	"fmt"
	"os"
)

func GenArr() ([3]any, [5]any) {
	var arr1 [3]any
	arr1[0] = "first string"
	arr1[1] = 2

	arrLiteral := [5]any{1, true, "string", 7.7, [2]int{10, 11}}

	return arr1, arrLiteral
}

func GenSlice() ([]any, []int, []any) {
	var slice1 []any
	slice1 = make([]any, 10)
	slice1[5] = 777

	sampleArr := [10]int{1, 2, 3, 4, 5}
	slice2 := sampleArr[3:8]
	slice2 = append(slice2, []int{999, 99999}...)

	sliceLiteral := []any{"a", "b"}

	return slice1, slice2, sliceLiteral
}

func GenMap() (map[any]any, map[any]any) {
	var map1 map[any]any
	map1 = make(map[any]any) // == map1 := make(map[any]any)
	map1["s"] = "super extream"
	map1[0] = "first elem"
	delete(map1, 0)

	mapLiteral := map[any]any{"test": 0, "is": 7, "cool": 77}
	val, ok := map1["non-exist"]
	fmt.Println("Try to get non-exist key's value =>", "value : ", val, ", IsSuccess : ", ok)

	return map1, mapLiteral
}

type StructSample struct {
	Elem1 string
	elem2 int
	elem3 bool
}

func (s StructSample) ShowMe() string {
	return fmt.Sprintf("Hi, %#v", s)
}

func GenStruct() (any, StructSample) {
	var struct1 struct {
		num  float64
		word string
	}
	struct1.num = 2.3
	struct1.word = "basic struct"

	customStruct := new(StructSample)
	customStruct2 := StructSample{Elem1: "literal struct", elem2: 10}
	return customStruct, customStruct2
}

func CliArgs() []string {
	return os.Args[1:]
}

func PrintSlice(s []int) string {
	return fmt.Sprintf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
