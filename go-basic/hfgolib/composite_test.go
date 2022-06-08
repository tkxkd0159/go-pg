package hfgolib_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/tkxkd0159/go-pg/go-basic/hfgolib"
	"reflect"
	"testing"
)

func TestPrintSlice(t *testing.T) {
	got := hfgolib.PrintSlice([]int{1, 2, 3, 4, 5})
	exp := "len=5 cap=5 [1 2 3 4 5]\n"
	if got != exp {
		t.Error("TestPrintSlice is failed")
	}
}

type StructSample struct {
	Elem1 string
	Elem2 int
	Elem3 bool
}

func TestGenStruct(t *testing.T) {
	got1, got2 := hfgolib.GenStruct()

	testCase := []any{
		&struct {
			Elem1 string
			Elem2 int
			Elem3 bool
		}{
			Elem1: "",
			Elem2: 0,
			Elem3: false,
		},
		&StructSample{Elem1: "", Elem2: 0, Elem3: false},
	}

	exp1 := &struct {
		Elem1 string
		Elem2 int
		Elem3 bool
	}{
		Elem1: "",
		Elem2: 0,
		Elem3: false,
	}

	for _, tc := range testCase {
		if !cmp.Equal(exp1, tc) {
			t.Logf("Diff between origin and target : %s", cmp.Diff(exp1, tc))
		}
	}

	if got1.(*hfgolib.StructSample).Elem1 != exp1.Elem1 {
		t.Error("Elem1 field value is different")
	}

	answer := [][3]string{
		{"Elem1", "string"},
		{"Elem2", "int"},
		{"Elem3", "bool"},
	}

	e := reflect.ValueOf(&got2).Elem()
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type
		if varName != answer[i][0] || varType.String() != answer[i][1] {
			t.Error("reflect test is failed")
		}
	}
}

func TestSetPrivateField(t *testing.T) {
	got := hfgolib.PrivateFields{}
	got.SetElem1(10)
	if got.Elem1() != 10 {
		t.Error("You didn't set elem1 ")
	}
}
