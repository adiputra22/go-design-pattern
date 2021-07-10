package abstract_factory

type Pricey interface {
	Price() float64
}

type Chair interface {
	Pricey
	IsIotEnabled() bool
	IsSoft() bool
}

type Dimention struct {
	Length, Width, Height int
}

type CoffeeTable interface {
	Pricey
	Size() Dimention
	IsFoldable() bool
}

type SofaStyle string

const (
	EuropeanStyle SofaStyle = "european"
	AmericanStyle SofaStyle = "american"
)

type Sofa interface {
	Pricey
	Style() SofaStyle
}

type FurnitureFactory interface {
	CreateChair() Chair
	CreateCoffeTable() CoffeeTable
	CreateSofa() Sofa
}
