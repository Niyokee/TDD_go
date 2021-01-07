package TDD

type Money struct {
	amount int
}
type Dollar struct {
	Money
}

func (m *Money) times(multiplier int) Money {
	return Money{amount: m.amount * multiplier}
}

func (m Money) equals(m1 Money) bool {
	return m == m1
}

type Franc struct {
	Money
}

func (f *Franc) times(multiplier int) Franc {
	return Franc{Money{amount: f.amount * multiplier}}
}
