package informa

import (
	"github.com/adiputra22/go-design-pattern/abstract_factory"
	"github.com/adiputra22/go-design-pattern/abstract_factory/informa/product"
)

// factory
type Informa struct {
}

func (i *Informa) CreateChair() abstract_factory.Chair {
	return &product.BeanBag{}
}

func (i *Informa) CreateCoffeTable() abstract_factory.CoffeeTable {
	return nil
}

func (i *Informa) CreateSofa() abstract_factory.Sofa {
	return nil
}
