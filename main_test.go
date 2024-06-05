package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	rootTests := []struct{
		equation QuadraticEquation
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