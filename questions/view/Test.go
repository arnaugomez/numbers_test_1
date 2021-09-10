package questions_view

import (
	"fmt"

	"github.com/arnaugomez/numbers_test_1/common/util"
	questions_domain "github.com/arnaugomez/numbers_test_1/questions/domain"
)

func greetOrCorrect(res questions_domain.Result) {
	if res.Ok {
		fmt.Println("Hooray! Correct answer.")
		fmt.Println("You win", res.Score, "points.")
	} else {
		fmt.Println("Your answer:", res.Answer, ". Correct answer:", res.Question.Ans())
		fmt.Println("You lose", res.Score*-1, "points.")
	}
}

func Test(questions []questions_domain.Question) []questions_domain.Result {
	var results = make([]questions_domain.Result, len(questions))
	for i, q := range questions {
		ans := util.ReadInt("What is ", q.FirstNumber, " " + q.Operation.String() + " ", q.SecondNumber, "?")
		res := questions_domain.GetResult(q, ans)
		greetOrCorrect(res)
		results[i] = res
	}
	return results
}
