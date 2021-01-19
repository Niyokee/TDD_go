package chapter1

type Bank struct {
}

func (b *Bank) reduce(source Expression, to string) *Money {
	return NewDollar(10)
}

func NewBank() *Bank {
	return &Bank{}
}
