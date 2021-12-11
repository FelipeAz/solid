package category

const (
	HotBeverages = "Hot Beverages"
)

type SubCategoryInterface interface {
	GetTax() float64
}

type SubCategory struct {
	Name   string
	Parent Category
}

func NewSubCategory(name string, parent Category) SubCategoryInterface {
	return SubCategory{
		Name:   name,
		Parent: parent,
	}
}

func (c SubCategory) GetTax() float64 {
	return c.Parent.Tax
}
