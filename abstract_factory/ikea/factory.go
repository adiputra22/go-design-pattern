package ikea

import (
	"github.com/adiputra22/go-design-pattern/abstract_factory"
	"github.com/adiputra22/go-design-pattern/abstract_factory/ikea/product"
)

// factory
type Ikea struct {
}

func (i *Ikea) CreateChair() abstract_factory.Chair {
	return &product.Leifarne{}
}

func (i *Ikea) CreateCoffeTable() abstract_factory.CoffeeTable {
	return &product.Vittsjo{}
}

func (i *Ikea) CreateSofa() abstract_factory.Sofa {
	return &product.Hemlingbe{}
}
