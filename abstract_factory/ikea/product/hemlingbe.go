package product

import "github.com/adiputra22/go-design-pattern/abstract_factory"

type Hemlingbe struct{}

func (h *Hemlingbe) Price() float64 {
	return 890000
}

func (h *Hemlingbe) Style() abstract_factory.SofaStyle {
	return abstract_factory.AmericanStyle
}
