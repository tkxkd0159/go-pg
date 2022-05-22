package hfgolib

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

func InputTerm(r io.Reader) string {
	fmt.Print("Enter a string : ")
	reader := bufio.NewReader(r)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return input
}
