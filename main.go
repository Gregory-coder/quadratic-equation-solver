package equation

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

var (
	ErrNoRealRoots    = errors.New("no real roots found")
	ErrInvalidFactors = errors.New("invalid factors")
	ErrInvalidSyntax  = errors.New("invalid syntax")
)

type QuadraticEquation struct {
	a, b, c float64
}

func Parse(input string) (*QuadraticEquation, error) {
	equation := QuadraticEquation{0, 0, 0}
	var err error

	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")

	parseRatios(input, &equation)
	return &equation, err
}

func parseRatios(input string, equation *QuadraticEquation) {
	isRightPart := false
	startIndex := 0
	for i := 1; i < len(input); i++ {
		if input[i] == '+' || input[i] == '-' || input[i] == '=' || i == len(input)-1 {
			var part string
			if i < len(input)-1 {
				part = input[startIndex:i]
			} else {
				part = input[startIndex : i+1]
			}
			k, category, err := parsePart(part)

			if err != nil {
				break
			}

			if isRightPart {
				k = -k
			}
			switch category {
			case 2:
				equation.a += k
			case 1:
				equation.b += k
			case 0:
				equation.c += k
			}
			startIndex = i
			if input[i] == '=' {
				startIndex++
				isRightPart = true
			}
		}
	}
}
func parsePart(str string) (k float64, category int, err error) {
	xPosition := strings.Index(str, "x")
	if xPosition == -1 {
		category = 0
	} else {
		degreePosition := strings.Index(str, "^")
		if degreePosition == -1 {
			category = 1
		} else {
			category = 2
		}
		str = str[:xPosition]
	}
	if len(str) == 0 {
		k = 1
	} else {
		k, err = strconv.ParseFloat(strings.TrimSpace(str), 64)
	}
	return
}

func (qe QuadraticEquation) Solve() ([2]float64, error) {
	// the equality 0*x^2 + 0*x + 0 = 0 has an infinite number of roots
	if qe.a == 0 && qe.b == 0 && qe.c == 0 {
		return [2]float64{}, ErrInvalidFactors
	}

	discriminant := math.Pow(qe.b, 2) - 4*qe.a*qe.c
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
