package calculator_test

import (
	"calculator"
	"testing"
	"math"
)

/*func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Add(2,2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}*/

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a,b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}

	for  _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a,b float64
		want float64
	}

	testCases := []testCase {
		{a: 3, b: 4, want: -1},
		{a: 5, b: 2, want: 3},
		{a: 6, b: 2, want: 4},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}
/*func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(4,2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}*/

func TestMultiply(t *testing.T) {
	t.Parallel()
	var want float64 = 9
	got := calculator.Multiply(3,3)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func closeEnough(a,b,tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a,b float64
		want float64
	}

	testCases := [] testCase {
		{a: 1, b: 3, want: 0.333333,},
		{a: -1, b: -1, want: 1},
		{a: 10, b: 2, want: 5},
	}

	for _,tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)

		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}

		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Divide(%f, %f) : want %f, got %f", tc.a,tc.b, tc.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a float64
		want float64
	}

	testCases := [] testCase {
		{a: 4, want: 2},
		{a: 9, want: 3},
		{a: 16, want: 4},
	}

	for _,tc := range testCases {
		got, err := calculator.Sqrt(tc.a)

		if err != nil {
			t.Fatalf("want no error for input, get %v", err)
		}

		if !closeEnough(tc.a,got, 1) {
			t.Errorf("Sqrt(%f): want %f, got %f", tc.a, tc.want, got)
		}
	}
}
