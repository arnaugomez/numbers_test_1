package questions_domain

type Result struct {
	Question *Question
	Answer   int
	Ok       bool
	Score    int
}
