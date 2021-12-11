package product

import "github.com/FelipeAz/solid/internal/app/liskov-substitution/category"

type ProductInterface interface {
	GetPriceWithTax(tax float64) float64
}

type Product struct {
	Name     string
	Category category.CategoryInterface
	Price    float64
	Amount   int
}

func NewProduct(name string, category category.CategoryInterface, price float64, amount int) Product {
	return Product{
		Name:     name,
		Category: category,
		Price:    price,
		Amount:   amount,
	}
}

func (p Product) GetPriceWithTax(tax float64) float64 {
	return (p.Price * (tax + p.Category.GetTax())) * float64(p.Amount)
}
