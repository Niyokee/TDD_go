package chapter1

type Dollar struct {
	amount int
}

// value-objectとして扱いたいので、新しいオブジェクトを返すようにする
func (d *Dollar) times(multiplier int) *Dollar {
	return NewDollar(d.amount * multiplier)
}

func (d *Dollar) equals(d1 *Dollar) bool {
	return d.amount == d1.amount
}

func NewDollar(amount int) *Dollar {
	return &Dollar{amount: amount}
}
