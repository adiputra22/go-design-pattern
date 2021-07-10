package product

type Leifarne struct{}

func (l *Leifarne) IsIotEnabled() bool {
	return false
}

func (l *Leifarne) IsSoft() bool {
	return false
}

func (v *Leifarne) Price() float64 {
	return 890000
}
