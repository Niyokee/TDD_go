package chapter1

import "testing"

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	five.times(2)
	if 10 != five.amount {
		t.Errorf("Dollar(5) times 2 is %v want 10", five.amount)
	}
}
