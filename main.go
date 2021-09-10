package main

import (
	"fmt"

	questions_view "github.com/arnaugomez/numbers_test_1/questions/view"
)

func main() {
	for {
		fmt.Println("Welcome to the go quiz by Arnau Gómez.")
		questions := questions_view.Setup()
		if(len(questions) == 0) {
			break
		}
		results := questions_view.Test(questions)
		questions_view.Results(results)
	}
}
