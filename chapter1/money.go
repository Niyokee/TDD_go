package chapter1

// DollarとFrancには共通部分がある。
// コピペで犯してきた罪を償う。
type Money struct {
	// DollarもFrancもamount を持つので移動させる
	amount int
}

// equal methodもDollarとFrancで共通しているので Moneyに引き上げた
// mとm1のamountが等価であるかどうか判定するメソッド
// 引数genericsにできないかな、、、
func (m *Money) equals(m1 *Money) bool {
	return m.amount == m1.amount
}

func NewMoney(amount int) *Money {
	return &Money{amount: amount}
}
