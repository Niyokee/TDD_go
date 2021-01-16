package chapter1

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.times(2)
	if 10 != product.amount {
		t.Errorf("Dollar(5) times 2 is %v want 10", product.amount)
	}

	product = five.times(3)
	if 15 != product.amount {
		t.Errorf("Dollar(5) times 3 is %v want 10", product.amount)
	}
}

// generics使えばきれいにかけそう
func TestEquality(t *testing.T) {
	if !NewDollar(5).equals(NewDollar(5)) {
		t.Errorf("Dollar(5) is not equal to %v want true", NewDollar(5))
	}
	if NewDollar(5).equals(NewDollar(6)) {
		t.Errorf("Dollar(5) is not equal to %v want true", NewDollar(6))
	}

	if !NewFranc(5).equals(NewFranc(5)) {
		t.Errorf("Franc(5) is not equal to %v want true", NewFranc(5))
	}
	if NewFranc(5).equals(NewFranc(6)) {
		t.Errorf("Franc(5) is not equal to %v want true", NewFranc(6))
	}
	//FrancとDollar比べたらどうなる？
	//FrancとDollarが等しくないことを確認する →　等しい！
	if NewFranc(5).equals(NewDollar(5)) {
		t.Errorf("Franc(5) is not equal to %v: type: %T want false", NewDollar(5), NewDollar(5), )
	}

}

// Dollar + Franc を実現するために、
// FrancでもDollarと同じことができるようにしよう。
func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	product := five.times(2)
	if 10 != product.amount {
		t.Errorf("Franc(5) times 2 is %v want 10", product.amount)
	}

	product = five.times(3)
	if 15 != product.amount {
		t.Errorf("Franc(5) times 3 is %v want 10", product.amount)
	}
}


func TestCurrency (t *testing.T) {
	// Dollar構造体の通貨がUSDであることを確かめる
	if NewDollar(1).currency != "USD" {
		t.Errorf("the currency of Dollar(1) is %v want USD", NewDollar(1).currency)
	}
	// Franc構造体の通貨がCHFであることを確かめる
	if NewFranc(1).currency != "CHF" {
		t.Errorf("the currency of Franc(1) is %v want CHF", NewFranc(1).currency)
	}
}
