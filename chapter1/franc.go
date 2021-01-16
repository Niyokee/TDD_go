package chapter1

// TDDのステップに則って、まずは仮実装でテストを全てグリーンにする
// とりあえずまるっとコピーしたクソコードでいい。
type Franc struct {
	*Money
}

// value-objectとして扱いたいので、新しいオブジェクトを返すようにする
func (f *Franc) times(multiplier int) *Franc {
	return NewFranc(f.amount * multiplier)
}

func NewFranc(amount int) *Franc {
	return &Franc{NewMoney(amount)}
}
