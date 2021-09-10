package questions_data

import (
	"github.com/arnaugomez/numbers_test_1/common/util"
	questions_domain "github.com/arnaugomez/numbers_test_1/questions/domain"
)

func GetQuestions(amount int) []questions_domain.Question {
	questions := make([]questions_domain.Question, 0)
	for amount > 0 {
		questions = append(questions, questions_domain.Question{
			FirstNumber:  util.RandInt(100),
			SecondNumber: util.RandInt(100),
			Operation:    questions_domain.RandOperation(),
		})
		amount = amount - 1
	}
	return questions
}

