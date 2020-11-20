package sports

type adidas struct {
}

func (a *adidas) makeShoe() iShoe {
	return &adidasShoe{
		shoe{
			logo: "adidas",
			size: 13,
		},
	}
}

func (a *adidas) makeShirt() iShirt {
	return &adidasShirt{
		shirt{
		logo: "adidas",
		size: 13,
	}}
}

type adidasShirt struct {
	shirt
}

type adidasShoe struct {
	shoe
}