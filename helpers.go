package equation

import (
	"strconv"
	"strings"
)

func parseRatios(input string, eq *QuadraticEquation) (err error) {
	isRightPart := false
	startIndex := 0
	// the string is split into fragments, which are parsed independently
	for i := 1; i < len(input); i++ {
		if input[i] == '+' || input[i] == '-' || input[i] == '=' || i == len(input)-1 {
			var part string
			if i < len(input)-1 {
				part = input[startIndex:i]
			} else {
				part = input[startIndex : i+1]
			}

			err = parseFragment(part, eq, isRightPart)

			if err != nil {
				return ErrInvalidSyntax
			}

			startIndex = i
			if input[i] == '=' {
				isRightPart = true
				startIndex++
			}
		}
	}
	return
}

func parseFragment(str string, eq *QuadraticEquation, isRightPart bool) (err error) {
	var xPosition, degreePosition int
	var coefficient float64
	// find 'x' (if it is here) and a symbol of squaring 'x' ('^')
	xPosition = strings.Index(str, "x")
	if xPosition != -1 {
		degreePosition = strings.Index(str, "^")
		str = str[:xPosition]
	}
	// check if there is a coefficient before x
	if len(str) == 0 {
		coefficient = 1
	} else {
		coefficient, err = strconv.ParseFloat(strings.TrimSpace(str), 64)
		if err != nil {
			return
		}
	}
	// if this component on the right of '=', we need to negate it
	if isRightPart {
		coefficient = -coefficient
	}
	// adding coefficient
	if xPosition == -1 {
		eq.c += coefficient
	} else if degreePosition == -1 {
		eq.b += coefficient
	} else {
		eq.a += coefficient
	}

	return
}
