package hfgolib

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func InputTerm() string {
	fmt.Print("Enter a string : ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return input
}
