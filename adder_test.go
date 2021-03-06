package adder

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		n1   int
		n2   int
		want int
	}{
		{n1: 0, n2: 0, want: 0},
		{n1: 1, n2: 0, want: 1},
		{n1: -1, n2: 0, want: -1},
		{n1: 3, n2: -3, want: 0},
		{n1: 5, n2: 16, want: 21},
		{n1: 694, n2: 1, want: 695},
	}

	for i, tt := range tests {
		t.Logf("[%02d] test %d + %d", i, tt.n1, tt.n2)
		sum := Add(tt.n1, tt.n2)
		if sum != tt.want {
			t.Errorf("unexpected value:\n- want: %v\n- got: %v", tt.want, sum)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		n1   int
		n2   int
		want int
	}{
		{n1: 0, n2: 0, want: 0},
		{n1: 1, n2: 0, want: 0},
		{n1: -1, n2: 0, want: 0},
		{n1: 3, n2: -3, want: -9},
		{n1: 5, n2: 16, want: 80},
		{n1: 694, n2: 1, want: 694},
	}

	for i, tt := range tests {
		t.Logf("[%02d] test %d * %d", i, tt.n1, tt.n2)
		product := Multiply(tt.n1, tt.n2)
		if product != tt.want {
			t.Errorf("unexpected value:\n- want: %v\n- got: %v", tt.want, product)
		}
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		n1   int
		n2   int
		want int
	}{
		{n1: 3, n2: 0, want: 0},
		{n1: 3, n2: 3, want: 1},
		{n1: -1, n2: 1, want: -1},
		{n1: 900, n2: 4, want: 225},
		{n1: 3, n2: 2, want: 1},
		{n1: 694, n2: 1, want: 694},
	}

	for i, tt := range tests {
		t.Logf("[%02d] test %d / %d", i, tt.n1, tt.n2)
		quotient := Divide(tt.n1, tt.n2)
		if quotient != tt.want {
			t.Errorf("unexpected value:\n- want: %v\n- got: %v", tt.want, quotient)
		}
	}
}
