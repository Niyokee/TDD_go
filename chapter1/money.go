package chapter1

// DollarとFrancには共通部分がある。
// コピペで犯してきた罪を償う。
type Money struct {
	// DollarもFrancもamount を持つので移動させる
	amount int
	// currencyが重複していたので、Moneyへ移動
	currency string
}

// equal methodもDollarとFrancで共通しているので Moneyに引き上げた
// mとm1のamountが等価であるかどうか判定するメソッド
// 引数genericsにできないかな、、、
func (m *Money) equals(m1 *Money) bool {
	return m.amount == m1.amount
}

// value-objectとして扱いたいので、新しいオブジェクトを返すようにする
// サブクラスで（goだから厳密にはサフクラスではないが便宜上。。。）完全に一致したメソッドができたので
//　親クラスに引き揚げた
func (m *Money) times(multiplier int) *Money {
	return NewMoney(m.amount * multiplier, m.currency)
}

func NewMoney(amount int, currency string) *Money {
	return &Money{amount: amount, currency: currency}
}
