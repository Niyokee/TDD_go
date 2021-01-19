package chapter1

type Sum struct {
	augend Money
	addend Money
}

func NewSum(augend, addend Money) Expression {
	return &Sum{augend: augend, addend: addend}
}
