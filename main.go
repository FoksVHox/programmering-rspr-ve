package main

import "fmt"

var (
	s, p = 0, 0
)

func main() {
	answer := askQuestion()
	if answer {

	}

	if s > 3 {
		fmt.Println("You have answered", s, "questions correctly.")
		return
	}
	main()
}

// make a function that will ask the user a question and return the answer
func askQuestion() bool {

	// make a slice of strings that will hold yes or no questions
	var (
		q = []string{
			"Do you want to learn Go?",
			"Do you want to learn Python?",
		}
	)

	// make a for loop that will iterate through the slice of strings and ask the user a yes or no question
	for _, v := range q {
		fmt.Println(v)
		scanln, _ := fmt.Scanln(&v)

		if scanln == 1 {
			return true
		}
	}
	return false
}
