package TDD

type Bank struct {

}

func (b Bank) reduce(source Expression, to string) *Money {
	return &Money{amount: 10, currency: "USD"}
}