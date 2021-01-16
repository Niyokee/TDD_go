package chapter1

// TDDのステップに則って、まずは仮実装でテストを全てグリーンにする
// とりあえずまるっとコピーしたクソコードでいい。
type Franc struct {
	*Money
}

func NewFranc(amount int) *Money {
	return NewMoney(amount, "CHF")
}
