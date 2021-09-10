package questions_domain

type Question struct {
	FirstNumber  int
	SecondNumber int
	Operation    Operation
}

func (q Question) Ans() int {
	switch q.Operation {
	case Sum:
		return q.FirstNumber + q.SecondNumber
	case Substraction:
		return q.FirstNumber - q.SecondNumber
	}
	return q.FirstNumber * q.SecondNumber
}

func (q Question) IsOk(ans int) bool {
	return ans == q.Ans()
}
