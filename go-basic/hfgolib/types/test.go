package types

import "testing"

type TestSuite interface {
	BeforeAll(tb testing.TB) []any
	BeforeEach(tb testing.TB) []any
	AfterEach(tb testing.TB) []any
	AfterAll(tb testing.TB) []any
}

type anyFunc func() error

func BeforeAll(tb testing.TB, cb anyFunc) {
	err := cb()
	if err != nil {
		tb.Fatal(err)
	}
}

func AfterAll(tb testing.TB, cb anyFunc) {
	err := cb()
	if err != nil {
		tb.Fatal(err)
	}
}
