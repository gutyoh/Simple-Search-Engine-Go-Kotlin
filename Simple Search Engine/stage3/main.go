package main

/*
[Simple Search Engine - Stage 3/6: String theory](https://hyperskill.org/projects/89/stages/496/implement)
-------------------------------------------------------------------------------
[Functional decomposition](https://hyperskill.org/learn/step/17506)
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	NumberOfPeoplePrompt      = "Enter the number of people:"
	AllPeoplePrompt           = "Enter all people:"
	NameOrEmailToSearchPrompt = "\nEnter a name or email to search all suitable people."

	MenuMsg                  = "\n=== Menu ==="
	ListOfPeopleMsg          = "\n=== List of people ==="
	NoMatchingPeopleFoundMsg = "No matching people found."
	ByeMsg                   = "\nBye!"
	IncorrectOptionMsg       = "\nIncorrect option! Try again."
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	people := readPeople(scanner)
	for {
		switch showMenu() {
		case 1:
			searchPeople(scanner, people)
		case 2:
			printPeople(people)
		case 0:
			fmt.Println(ByeMsg)
			return
		default:
			fmt.Println(IncorrectOptionMsg)
		}
	}
}

func showMenu() int {
	fmt.Println(MenuMsg)
	fmt.Println("1. Find a person")
	fmt.Println("2. Print all people")
	fmt.Println("0. Exit")

	var choice int
	fmt.Scanln(&choice)

	return choice
}

func readPeople(scanner *bufio.Scanner) []string {
	fmt.Println(NumberOfPeoplePrompt)
	var numberOfPeople int
	fmt.Scanln(&numberOfPeople)

	fmt.Println(AllPeoplePrompt)
	people := make([]string, 0, numberOfPeople)
	for i := 0; i < numberOfPeople; i++ {
		scanner.Scan()
		people = append(people, scanner.Text())
	}
	return people
}

func searchPeople(scanner *bufio.Scanner, people []string) {
	fmt.Println(NameOrEmailToSearchPrompt)
	scanner.Scan()
	query := strings.ToLower(scanner.Text())
	found := false
	for _, person := range people {
		if strings.Contains(strings.ToLower(person), query) {
			fmt.Println(person)
			found = true
		}
	}
	if !found {
		fmt.Println(NoMatchingPeopleFoundMsg)
	}
}

func printPeople(people []string) {
	fmt.Println(ListOfPeopleMsg)
	for _, person := range people {
		fmt.Println(person)
	}
}
