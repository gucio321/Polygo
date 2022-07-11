package polynomials

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// ErrUnexpectedEquation is returned by Parse if there is something wrong with equation string given
// you can refer to this error by calling errors.Is(err, ErrUnexpectedEquation) in order to detect it.
var ErrUnexpectedEquation = errors.New("unexpected equation")

type Equation struct {
	MonomialsGroups []*MonomialsGroup
}

func Parse(equation string) (result *Equation, err error) {
	result = &Equation{}
	// remove all spaces - they are user-friendly but they are not necessary
	equation = strings.ReplaceAll(equation, " ", "")

	// split to right and left sides
	splited := strings.Split(equation, "=")

	var leftSideStr, rightSideStr string

	switch len(splited) {
	// only left side - x^2 + 2x = 0
	case 1:
		leftSideStr = splited[0]
	// standard equation x^2 + 2x = 8x + 2
	case 2:
		leftSideStr, rightSideStr = splited[0], splited[1]
	// wtf???
	default:
		return nil, ErrUnexpectedEquation
	}

	leftSide, rightSide := parseStr(leftSideStr), parseStr(rightSideStr)

	// move everything from right to left side

	for _, x := range rightSide {
		x.ChangeSide()
		leftSide = append(leftSide, x)
	}

	result.AddGroup(newGroup().Add(leftSide...))
	result.Simplify()
	return result, nil
}

func parseStr(s string) []*Monomial {
	result := make([]*Monomial, 0)
	if len(s) == 0 {
		return result
	}

	if s[0] != '-' {
		result = append(result, newMonomial())
	}

	for i := 0; i < len(s); i++ {
		currentLetter := s[i]
		switch {
		case currentLetter == '-':
			m := newMonomial()
			m.Coefficient *= -1
			result = append(result, m)
		case currentLetter == '+':
			result = append(result, newMonomial())
		case currentLetter == '^':
			last := result[len(result)-1]
			p := "0"
			for j := i + 1; j < len(s) && ((s[j] >= '0' && s[j] <= '9') || s[j] == '.'); j++ {
				p += string(s[j])
			}

			power, err := strconv.ParseFloat(p, 64)
			if err != nil {
				log.Panicf("Unexpected error occured - this couldn't have happened: %v", err)
			}

			last.Power = power

			// -1 for base 0 in p value
			i += len(p) - 1
		case (currentLetter >= '0' && currentLetter <= '9') || currentLetter == '.':
			last := result[len(result)-1]
			c := "0"
			for j := i; j < len(s) && ((s[j] >= '0' && s[j] <= '9') || s[j] == '.'); j++ {
				c += string(s[j])
			}

			coefficient, err := strconv.ParseFloat(c, 64)
			if err != nil {
				log.Panicf("Unexpected error occured - this couldn't have happend: %v", err)
			}

			last.Coefficient *= coefficient

			// - 1 for base 0 value of c and -1 for i++ in for assignment
			i += len(c) - 2
		default:
			last := result[len(result)-1]
			last.Power = 1
		}
	}
	return result
}

// AddGroup adds a new monomials group
func (e *Equation) AddGroup(g *MonomialsGroup) {
	e.MonomialsGroups = append(e.MonomialsGroups, g)
}

// Simplify simplifies each of monomials group
func (e *Equation) Simplify() {
	for _, g := range e.MonomialsGroups {
		g.Simplify()
	}
}

// String implements fmt.Stringer
func (e *Equation) String() string {
	result := ""
	for _, i := range e.MonomialsGroups {
		result += i.String()
	}

	return fmt.Sprintf("%v = 0", result)
}
