package category

const (
	Beverage = "Beverage"
	Meat     = "Meat"
)

type CategoryInterface interface {
	GetTax() float64
}

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

func (c Category) GetTax() float64 {
	return c.Tax
}
