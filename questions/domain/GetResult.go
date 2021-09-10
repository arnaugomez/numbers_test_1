package questions_domain

import "github.com/arnaugomez/numbers_test_1/common/util"

func GetResult(q Question, ans int) Result {
	isOk := q.IsOk(ans)
	sign := -1
	if isOk {
		sign = 1
	}
	return Result{
		Question: &q,
		Answer: ans,
		Ok: isOk,
		Score: util.RandInt(10) * sign,
	}
}