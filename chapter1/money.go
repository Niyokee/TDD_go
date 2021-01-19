package chapter1

// Money is aDollarとFrancには共通部分がある。
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
	return *m == *m1
}

// value-objectとして扱いたいので、新しいオブジェクトを返すようにする
// サブクラスで（goだから厳密にはサフクラスではないが便宜上。。。）完全に一致したメソッドができたので
//　親クラスに引き揚げた
func (m *Money) times(multiplier int) *Money {
	return NewMoney(m.amount*multiplier, m.currency)
}

func (m *Money) plus(m1 *Money) Expression {
	return NewSum(*m, *m1)
}

func NewMoney(amount int, currency string) *Money {
	return &Money{amount: amount, currency: currency}
}

func NewDollar(amount int) *Money {
	return NewMoney(amount, "USD")
}

func NewFranc(amount int) *Money {
	return NewMoney(amount, "CHF")
}
