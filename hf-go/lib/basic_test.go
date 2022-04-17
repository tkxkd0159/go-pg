package lib

import (
	"testing"
)

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
	ans2 := IntMin(-2, 2)
	if ans2 != -2 {
		t.Errorf("IntMin(-2, 2) = %d; want -2", ans2)
	}
}
