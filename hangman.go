package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func string_limiter(source_string string, limit int) []string {
	out := strings.Split(source_string, "")
	return out[:limit]
}

func main() {

	words := make(map[int]string)
	words[0] = "search"
	words[1] = "banana"
	words[2] = "orange"
	words[3] = "arguments"

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	var input string

	lives := 5
  counter := make([]string, 0)
  
	fmt.Println("Welcome to Hangman. Guess a letter from the word:")
	for {
		if lives == 0 {
			fmt.Println("You have no more lives left. Goodbye!")
			break
		}

    // take input from the user, obtain the first char
		fmt.Scanf("%s", &input)
		first_char := string_limiter(input, 1)

    
    // check if char is contained within word, decrease lives
    // if incorrect char
		if strings.Contains(words[0], first_char[0]) {
			fmt.Println("You guessed correctly")
      
		} else {
			lives--
			fmt.Printf("You lose a life: Lives: %d \n", lives)
      
		}
    counter = append(counter, first_char[0])
    fmt.Println(counter)
    
	}
