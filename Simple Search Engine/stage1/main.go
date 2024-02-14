package main

/*
[Simple Search Engine - Stage 1/6: String theory](https://hyperskill.org/projects/89/stages/494/implement)
-------------------------------------------------------------------------------
[Loops](https://hyperskill.org/learn/step/14707)
[Control statements](https://hyperskill.org/learn/step/16235)
[Advanced input](https://hyperskill.org/learn/step/18567)
[Operations with strings](https://hyperskill.org/learn/step/18548)
[Working with slices](https://hyperskill.org/learn/step/15935)
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Read the first line of words
	scanner.Scan()
	wordsLine := scanner.Text()
	words := strings.Split(wordsLine, " ")

	// Read the search word
	scanner.Scan()
	searchWord := scanner.Text()

	// Search for the word and output the result
	found := false
	for index, word := range words {
		if word == searchWord {
			fmt.Println(index + 1) // index + 1 because we start counting from 1
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Not found")
	}
}
