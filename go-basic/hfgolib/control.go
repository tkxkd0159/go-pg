package hfgolib

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func ForUnicode(word string) map[string]rune {
	wordMap := make(map[string]rune)
	for _, r := range word {
		wordMap[string(r)] = r
	}

	return wordMap
}

func GuessGame() bool {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("Random number will be generated between 1 and 100")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	totalGuess, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		log.Fatal(err)
	}
	for guesses := 0; guesses < totalGuess; guesses++ {
		fmt.Println("You have", totalGuess-guesses, "guesses left")
		input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Your guess was LOW")
		} else if guess > target {
			fmt.Println("Your guess was HIGH")
		} else {
			fmt.Println("Your guess was correct")
			return true
		}
	}
	fmt.Println("You use all guess chances")
	return false
}
