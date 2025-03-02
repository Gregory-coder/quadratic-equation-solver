package equation

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	got := New(0.0001, -7, 5)
	want := QuadraticEquation{0.0001, -7, 5}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("got %v, expected %v", got, want)
	}

}
func TestSolve(t *testing.T) {
	rootTests := []struct {
		equation      QuadraticEquation
		expectedValue [2]float64
		expectedError error
	}{
		{QuadraticEquation{1.0, -5.0, 6.0}, [2]float64{2.0, 3.0}, nil},
		{QuadraticEquation{1.0, -5.0, 15.0}, [2]float64{}, ErrNoRealRoots},
		{QuadraticEquation{}, [2]float64{}, ErrInvalidFactors},
	}

	for _, tt := range rootTests {
		got, err := tt.equation.Solve()
		if got != tt.expectedValue {
			t.Errorf("got %v, expected %v", got, tt.expectedValue)
		}
		if err != tt.expectedError {
			t.Fatal("got unexpected error: ", err)
		}
	}
}

func TestParse(t *testing.T) {
	parseTests := []struct {
		input            string
		expectedEquation QuadraticEquation
		expectedError    error
	}{
		{"x^2 = 4", QuadraticEquation{1, 0, -4}, nil},
		{"-5x + 2x + 7 = 0", QuadraticEquation{0, -3, 7}, nil},
		{"6x^2 = 10x", QuadraticEquation{6, -10, 0}, nil},
		{"6*x^2 = 10x", QuadraticEquation{0, 0, 0}, ErrInvalidSyntax},
		{"3.5x^2 - 2.0001x + 90 = 0 + 2x - 4", QuadraticEquation{3.5, -4.0001, 94}, nil},
		{" ", QuadraticEquation{0, 0, 0}, nil},
	}

	for _, tt := range parseTests {
		got, err := Parse(tt.input)
		if !reflect.DeepEqual(*got, tt.expectedEquation) {
			t.Errorf("got %v, expected %v", got, tt.expectedEquation)
		}
		if err != tt.expectedError {
			t.Fatal("got unexpected error: ", err)
		}
	}
}

func BenchmarkParse(b *testing.B) {
	testEquations := []string{
		" ",
		"x^2 = 4",
		"6x^2 = 10x",
		"6*x^2 = 10x",
		"-5x + 2x + 7 = 0",
		"3.5x^2 — 2.0001x + 90 = 0 + 2x - 4",
	}
	b.ResetTimer()
	for i := 0; i < 100; i++ {
		for _, tt := range testEquations {
			Parse(tt)
		}
	}
}
