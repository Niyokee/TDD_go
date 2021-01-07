package TDD

import (
	"reflect"
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := Dollar{Money{amount: 5}}
	ten := Dollar{Money{amount: 10}}
	fifteen := Dollar{Money{amount: 15}}
	if reflect.DeepEqual(five.times(2), ten) {
		t.Errorf("Dollor(5) times 2 = %d; want 10", ten)
	}
	if reflect.DeepEqual(five.times(3), fifteen) {
		t.Errorf("Dollor(5) times 3 = %d; want 15", fifteen)
	}
}

func TestEquality(t *testing.T) {
	d := Dollar{Money{amount: 5}}
	d1 :=  Dollar{Money{amount: 5}}
	d2 :=  Dollar{Money{amount: 6}}

	f := Franc{Money{amount: 5}}
	f1 :=  Franc{Money{amount: 5}}
	f2 :=  Franc{Money{amount: 6}}

	if !reflect.DeepEqual(d, d1) {
		t.Errorf("d: %d, d1: %d", d, d1)
	}
	if reflect.DeepEqual(d, d2) {
		t.Errorf("d: %d, d1: %d", d, d2)
	}

	if !reflect.DeepEqual(f, f1) {
		t.Errorf("f: %d, f1: %d", f, f1)
	}
	if reflect.DeepEqual(f, f2) {
		t.Errorf("f: %d, f1: %d", f, f2)
	}

}

func TestFrancMultiplication(t *testing.T) {
	five := Franc{Money{amount: 5}}
	ten := Franc{Money{amount: 10}}
	fifteen := Franc{Money{amount: 15}}

	if !reflect.DeepEqual(five.times(2), ten) {
		t.Errorf("Dollor(5) times 2 = %d; want 10", ten)
	}
	if !reflect.DeepEqual(five.times(3), fifteen) {
		t.Errorf("Dollor(5) times 3 = %d; want 15", fifteen)
	}
}
