package polynomials

import (
	"strings"
)

// MonomialsGroup represents a group of monomials (example: (x^2+4+3x)^4 )
type MonomialsGroup struct {
	Power     float64
	Monomials []*Monomial
}

func newGroup() *MonomialsGroup {
	return &MonomialsGroup{
		Power:     1,
		Monomials: make([]*Monomial, 0),
	}
}

// Add adds monomials to the group
func (g *MonomialsGroup) Add(m ...*Monomial) *MonomialsGroup {
	g.Monomials = append(g.Monomials, m...)
	return g
}

// Simplify simplifies and sorts group
func (g *MonomialsGroup) Simplify() {
	var maxPower float64 = 0

	// key is power, value is coefficient
	monomials := make(map[float64]float64)

	for _, m := range g.Monomials {
		monomials[m.Power] += m.Coefficient

		if m.Power > maxPower {
			maxPower = m.Power
		}
	}

	newMonomials := make([]*Monomial, 0)
	for i := maxPower; i >= 0; i-- {
		c, contains := monomials[i]
		if !contains || c == 0 {
			continue
		}
		newMonomials = append(newMonomials, &Monomial{
			Power:       i,
			Coefficient: c,
		})
	}

	g.Monomials = newMonomials
}

// String implements fmt.Stringer
func (g *MonomialsGroup) String() string {
	result := ""

	for _, i := range g.Monomials {
		result += i.String()
	}

	result = strings.TrimLeft(result, "+ -")
	result = strings.TrimRight(result, " ")

	if len(g.Monomials) > 1 {
		result = "(" + result + ")"
	}

	return result
}
