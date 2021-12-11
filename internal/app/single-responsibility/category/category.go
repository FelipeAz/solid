package category

const (
	Beverage = "Beverage"
	Meat     = "Meat"
)

type Category struct {
	Name string
	Tax  float64
}

func NewCategory(name string, tax float64) Category {
	return Category{
		Name: name,
		Tax:  tax,
	}
}
