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
	if !NewDollar(5).Money.equals(NewDollar(5).Money) {
		t.Errorf("Dollar(5) is not equal to %v want true", NewDollar(5))
	}
	if NewDollar(5).Money.equals(NewDollar(6).Money) {
		t.Errorf("Dollar(5) is not equal to %v want true", NewDollar(6))
	}

	if !NewFranc(5).Money.equals(NewFranc(5).Money) {
		t.Errorf("Franc(5) is not equal to %v want true", NewFranc(5))
	}
	if NewFranc(5).Money.equals(NewFranc(6).Money) {
		t.Errorf("Franc(5) is not equal to %v want true", NewFranc(6))
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
