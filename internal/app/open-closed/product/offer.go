package product

import "github.com/FelipeAz/solid/internal/app/open-closed/category"

type ProductOffer struct {
	Name     string
	Category category.CategoryInterface
	Price    float64
	Amount   int
	Offer    float64
}

func NewProductOffer(
	name string,
	category category.CategoryInterface,
	price float64,
	amount int,
	offer float64) ProductOffer {
	return ProductOffer{
		Name:     name,
		Category: category,
		Price:    price,
		Amount:   amount,
		Offer:    offer,
	}
}

func (p ProductOffer) getPriceWithOffer() float64 {
	return p.Price * (1 - p.Offer)
}

func (p ProductOffer) GetPriceWithTax(tax float64) float64 {
	return (p.getPriceWithOffer() * (tax + p.Category.GetTax())) * float64(p.Amount)
}
