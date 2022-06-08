package hfgolib_test

import (
	"fmt"
	h "github.com/tkxkd0159/go-pg/go-basic/hfgolib"
	"testing"
)

func TestHttpGet(t *testing.T) {
	c1 := make(chan h.Webpage)
	c2 := make(chan h.Webpage)
	chs := []chan h.Webpage{c1, c2}
	urls := []string{"https://example.com", "https://google.com"}
	for i, url := range urls {
		go h.HttpGet(url, chs[i])
	}
	page1 := <-c1
	page2 := <-c2
	t.SkipNow()

	if page2.ResSize() == 14893 {
		t.Error("ERROR: Your second HTTP GET is wrong", page2.ResSize())
		t.Log("LOG: Your second HTTP GET is wrong", page2.ResSize())
	}
	if page1.ResSize() != 1256 {
		t.Error("Your first HTTP GET is wrong")
	}

	if testing.Verbose() {
		fmt.Println("I set verbose flag")
	}
	if testing.Short() {
		fmt.Println("I set short flag")
	}
}
