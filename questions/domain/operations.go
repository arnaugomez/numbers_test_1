package questions_domain

import (
	"github.com/arnaugomez/numbers_test_1/common/util"
)

type Operation int

const (
	Sum Operation = iota + 1
	Substraction
	Multiplication
)

func RandOperation() Operation {
	return Operation(util.RandInt(3))
}

func (o Operation) String() string {
	switch o {
	case Sum:
		return "+"
	case Substraction:
		return "-"
	}
	return "*"
}
