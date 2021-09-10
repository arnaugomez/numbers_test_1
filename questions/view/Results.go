package questions_view

import (
	"fmt"

	questions_domain "github.com/arnaugomez/numbers_test_1/questions/domain"
)

type stats struct {
	score int
	ok int
}

func getStats(r []questions_domain.Result) stats {
	s := stats{}
	for _, r := range r {
		if(r.Ok) {
			s.ok = s.ok + 1
		}
		s.score = s.score + r.Score
	}
	return s
}

func Results(r []questions_domain.Result) {
	s := getStats(r)
	fmt.Println("")
	fmt.Println("--------------------------------")
	fmt.Println("------------Results-------------")
	fmt.Println("--------------------------------")
	fmt.Println("You answered correctly ", s.ok, " out of ", len(r), " questions.")
	fmt.Println("Total score: ", s.score)
}