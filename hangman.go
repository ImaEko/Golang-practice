package main

import (
	"fmt"
	"strings"
)

type letter struct {
	char     rune
	revealed bool
}

type word []letter

func string_limiter(source_string string, limit int) []string {
	out := strings.Split(source_string, "")
	return out[:limit]
}

// This will print out the current word
// Revealed letters will be shown as that letter. Unrevealed letters will be shown as unscroll
func stringHint(userWord word) string {
	var hint string
	for _, l := range userWord {
		if l.revealed {
			hint = fmt.Sprintf("%s %s", hint, string(l.char))
		} else {
			hint = fmt.Sprintf("%s _", hint)
		}
	}

	return hint[1:]
}

// check for matching letters
// if match then reveal letters
// if not the life lost
func guessedLetters(guess string, userWord word) (word, bool) {

	return userWord, true
}

func stringToWord(myWord string) word {
	convertedWord := make(word, 0)
	for _, char := range myWord {
		temp := letter{
			char:     char,
			revealed: false,
		}
		convertedWord = append(convertedWord, temp)
	}
	return convertedWord

}

func main() {
	test_word := "apple"
	conv_word := stringToWord(test_word)

	for _, l := range conv_word {
		fmt.Println(string(l.char))
	}

	/*
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

	   }*/
}
