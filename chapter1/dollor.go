package chapter1

type Dollar struct {
	*Money
}

// value-objectとして扱いたいので、新しいオブジェクトを返すようにする
func (d *Dollar) times(multiplier int) *Dollar {
	return NewDollar(d.amount * multiplier)
}

func NewDollar(amount int) *Dollar {
	return &Dollar{NewMoney(amount)}
}
