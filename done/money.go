package done

type Money struct {
	amount   int
	currency string
}

func NewMoney(amount int, currency string) *Money {
	m := new(Money)
	m.amount = amount
	m.currency = currency
	return m
}

func (m Money) equals(m1 Money) bool {
	return m == m1
}

func (m Money) times(multiplier int) *Money {
	return NewMoney(m.amount*multiplier, m.currency)
}

func (m *Money) plus(addend *Money) *Money {
	return NewMoney(m.amount + addend.amount, m.currency)
}

func NewDollar(amount int) *Money {
	return NewMoney(amount, "USD")
}

func NewFranc(amount int) *Money {
	return NewMoney(amount, "USD")
}
