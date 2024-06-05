package main

import (
	"errors"
	"math"
)

var ErrNoRealRoots = errors.New("no real roots found") 
var ErrInvalidFactors = errors.New("invalid factors") 

type QuadraticEquation struct {
	a, b, c float64
}

func (qe QuadraticEquation) Solve() ([2]float64, error) {
	// the equality 0*x^2 + 0*x + 0 = 0 has an infinite number of roots
	if (qe.a == 0 && qe.b == 0 && qe.c == 0) {
		return [2]float64{}, ErrInvalidFactors
	}

	discriminant := math.Pow(qe.b, 2) - 4 * qe.a * qe.c
	// the equality with such D does not have real roots
	if discriminant < 0 {
		return [2]float64{}, ErrNoRealRoots
	}

	roots := [2]float64{}
	roots[0] = (-qe.b - math.Sqrt(discriminant)) / (2 * qe.a)
	roots[1] = (-qe.b + math.Sqrt(discriminant)) / (2 * qe.a)

	// the roots are returned in increasing order
	return roots, nil
}