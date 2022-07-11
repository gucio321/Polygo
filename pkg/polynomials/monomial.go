package polynomials

import "fmt"

// Monomial represents a single monomial (example: 2x^2)
type Monomial struct {
	Positive    bool
	Coefficient float64
	Power       float64
}

func newMonomial() *Monomial {
	return &Monomial{
		Coefficient: 1,
		Power:       0,
	}
}

// String implements fmt.Stringer interface.
// it is like ToString in C# - makes fmt.Print and familiar to print
// this as human-readable
func (m *Monomial) String() string {
	var result string
	if m.Coefficient >= 0 {
		result += "+"
	}

	result += fmt.Sprintf("%v", m.Coefficient)

	if m.Power == 0 {
		return result
	}

	if m.Power > 0 {
		result = fmt.Sprintf("%sx^%v", result, m.Power)
	}

	return result
}

func (s *Monomial) ChangeSide() {
	s.Coefficient *= -1
}
