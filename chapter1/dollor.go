package chapter1

type Dollar struct {
	*Money
}

func NewDollar(amount int) *Money {
	return NewMoney(amount, "USD")
}
