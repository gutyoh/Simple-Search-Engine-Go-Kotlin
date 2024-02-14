package main

/*
[Simple Search Engine - Stage 2/6: String theory](https://hyperskill.org/projects/89/stages/495/implement)
-------------------------------------------------------------------------------
[String search](https://hyperskill.org/learn/step/18995)
[String formatting](https://hyperskill.org/learn/step/16860)
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	NumberOfPeoplePrompt     = "Enter the number of people:"
	AllPeoplePrompt          = "Enter all people:"
	NumberOfQueriesPrompt    = "\nEnter the number of search queries:"
	DataToSearchPeoplePrompt = "\nEnter data to search people:"

	PeopleFoundMsg           = "\nPeople found:"
	NoMatchingPeopleFoundMsg = "No matching people found."
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(NumberOfPeoplePrompt)
	var numberOfPeople int
	fmt.Scanln(&numberOfPeople)

	fmt.Println(AllPeoplePrompt)
	people := make([]string, 0, numberOfPeople)
	for i := 0; i < numberOfPeople; i++ {
		if scanner.Scan() {
			people = append(people, scanner.Text())
		}
	}

	fmt.Println(NumberOfQueriesPrompt)
	var numberOfQueries int
	fmt.Scanln(&numberOfQueries)

	for i := 0; i < numberOfQueries; i++ {
		fmt.Println(DataToSearchPeoplePrompt)
		if scanner.Scan() {
			query := strings.ToLower(scanner.Text())
			found := false
			fmt.Println(PeopleFoundMsg)
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
	}
}
