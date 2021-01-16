package chapter1

// TDDのステップに則って、まずは仮実装でテストを全てグリーンにする
// とりあえずまるっとコピーしたクソコードでいい。
type Franc struct {
	amount int
}

// value-objectとして扱いたいので、新しいオブジェクトを返すようにする
func (f *Franc) times(multiplier int) *Franc {
	return NewFranc(f.amount * multiplier)
}

func (f *Franc) equals(f1 *Franc) bool {
	return f.amount == f1.amount
}

func NewFranc(amount int) *Franc {
	return &Franc{amount: amount}
}

