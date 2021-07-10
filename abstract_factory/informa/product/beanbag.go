package product

type BeanBag struct{}

func (b *BeanBag) Price() float64 {
	return 290222
}

func (b *BeanBag) IsIotEnabled() bool {
	return false
}

func (b *BeanBag) IsSoft() bool {
	return true
}
