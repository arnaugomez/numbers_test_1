package questions_view

import (
	"github.com/arnaugomez/numbers_test_1/common/util"
	questions_data "github.com/arnaugomez/numbers_test_1/questions/data"
	questions_domain "github.com/arnaugomez/numbers_test_1/questions/domain"
)

func Setup() []questions_domain.Question {
	numQuestions := util.ReadInt("How many questions do you want? ")
	return questions_data.GetQuestions(numQuestions)
}
