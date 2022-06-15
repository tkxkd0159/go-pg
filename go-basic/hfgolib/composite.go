package hfgolib

import (
	"fmt"
	"os"
	"reflect"
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
	Elem2 int
	Elem3 bool
}

func (s *StructSample) SetElem1(e any) {
	if reflect.TypeOf(e).String() == "string" {
		s.Elem1 = e.(string)
	}

}

func (ss StructSample) SetE1(s string) StructSample {
	ss.Elem1 = s
	return ss
}

func (ss *StructSample) SetE1Wptr(s string) {
	ss.Elem1 = s
}

func (ss StructSample) SetE2(i int) StructSample {
	ss.elem2 = i
	return ss
}

func (ss *StructSample) SetE2Wptr(i int) {
	ss.elem2 = i
}

func (ss StructSample) ShowMe() string {
	return fmt.Sprintf("Hi, %#v", ss)
}

func GenStruct() (any, StructSample) {
	var struct1 struct {
		num  float64
		word string
	}
	struct1.num = 2.3
	struct1.word = "basic struct"

	customStruct := new(StructSample)
	customStruct2 := StructSample{Elem1: "literal struct", Elem2: 10}
	return customStruct, customStruct2
}

func CliArgs() []string {
	return os.Args[1:]
}

func PrintSlice(s []int) string {
	return fmt.Sprintf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

type PrivateFields struct {
	elem1 int
	elem2 int
	Elem3 int
}

func (ps *PrivateFields) SetElem1(num int) {
	ps.elem1 = num
}

func (ps *PrivateFields) Elem1() int {
	return ps.elem1
}

func (ps PrivateFields) SetElem3(num int) int {
	return ps.Elem3 + num
}

// ## interface ##

type Recorder struct {
	List    []string
	Battery int
	mic     bool
}

func (r *Recorder) ToggleMic() {
	if r.mic == false {
		r.mic = true
	} else {
		r.mic = false
	}
}

func (r Recorder) Play() {
	for _, rec := range r.List {
		fmt.Printf("%s is playing now", rec)
	}
}

func (r Recorder) Stop() {
	fmt.Println("Stop recoreder")
}

type Player struct {
	List    []string
	battery int
}

func (p Player) Play() {
	for _, song := range p.List {
		fmt.Printf("%s is playing now", song)
	}
}

func (p Player) Stop() {
	fmt.Println("Stop recoreder")
}

func (p *Player) ChargeBattery(volt int) {
	p.battery += volt
}

type MediaMachine interface {
	Play()
	Stop()
}

var (
	_ MediaMachine = Recorder{}
	_ MediaMachine = &Player{}
)
