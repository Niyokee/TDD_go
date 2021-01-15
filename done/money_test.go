package TDD

import (
	"testing"
)

func TestSimpleAdd(t *testing.T)  {
	five := NewDollar(5)
	sum := five.plus(five)
	bank := Bank{}
	reduced := bank.reduce(sum, "USD")
	if *sum != *reduced {
		t.Errorf("Dollor(10) = %v; want 10", reduced)
	}
}
func TestMultiplication(t *testing.T) {
	five := NewDollar(5)

	if *five.times(2) != *NewDollar(10) {
		t.Errorf("Dollor(5) times 2 = %v; want 10", NewDollar(10))
	}
	if *five.times(3) != *NewDollar(15) {
		t.Errorf("Dollor(5) times 3 = %v; want 15", NewDollar(15))
	}
}

func TestEquality(t *testing.T) {
	d := NewDollar(5)
	d1 := NewDollar(5)
	d2 := NewDollar(6)

	f := NewFranc(5)

	if !d.equals(*d1) {
		t.Errorf("d: %v, d1: %v", d, d1)
	}
	if d1.equals(*d2) {
		t.Errorf("d: %v, d1: %v", d1, d2)
	}

	if !f.equals(*d) {
		t.Errorf("d: %v, d1: %v", f, d)
	}
}
