package hfgolib

import (
	"io"
	"log"
	"net/http"
)

type Webpage struct {
	Url  string
	Body string
}

type ResParser interface {
	ResSize(string) int
}

func HttpGet(url string, ch chan<- Webpage) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() // release the network connection
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	ch <- Webpage{
		url, string(body),
	}
}

func (p Webpage) ResSize() int {
	return len(p.Body)
}
