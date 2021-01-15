package chapter1

type Dollar struct {
	amount int
}

func (d *Dollar) times(multiplier int) {
	d.amount *= multiplier
}

func NewDollar(amount int) *Dollar {
	return &Dollar{amount: amount}
}
