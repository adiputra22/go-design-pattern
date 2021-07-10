package product

import "github.com/adiputra22/go-design-pattern/abstract_factory"

type Vittsjo struct{}

func (v *Vittsjo) Size() abstract_factory.Dimention {
	return abstract_factory.Dimention{
		Length: 20,
		Width:  30,
		Height: 10,
	}
}

func (v *Vittsjo) IsFoldable() bool {
	return false
}

func (v *Vittsjo) Price() float64 {
	return 890000
}
