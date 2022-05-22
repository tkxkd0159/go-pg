package hfgolib_test

import (
	"fmt"
	"github.com/tkxkd0159/go-pg/go-basic/hfgolib"
	"github.com/tkxkd0159/go-pg/go-basic/hfgolib/types"
	"os"
	"strings"
	"testing"
)

func TestInputTerm(t *testing.T) {
	expected := "Hello, Reader!\n"
	r := strings.NewReader(expected)
	got := hfgolib.InputTerm(r)
	if got != expected {
		t.Error("InputTerm is implemented wrongly")
	}
}

func TestRW(t *testing.T) {
	var f1 *os.File
	var err error
	types.BeforeAll(t,
		func() error {
			f1, err = os.Create("./f1.txt")
			if err != nil {
				return err
			}
			return nil
		},
	)
	defer types.AfterAll(t,
		func() error {
			err = f1.Close()
			err2 := os.Remove("./f1.txt")
			if (err != nil) || (err2 != nil) {
				return err
			}
			return nil
		},
	)
	_, _ = fmt.Fprintf(f1, "Data to write\n")

}
