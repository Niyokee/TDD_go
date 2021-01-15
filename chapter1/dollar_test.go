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

func TestEquality(t *testing.T)  {
	five := NewDollar(5)
	if !five.equals(NewDollar(5)) {
		t.Errorf("Dollar(5) is not equal to %v want true", NewDollar(5))
	}
	if five.equals(NewDollar(6)) {
		t.Errorf("Dollar(5) is not equal to %v want true", NewDollar(6))
	}
}
