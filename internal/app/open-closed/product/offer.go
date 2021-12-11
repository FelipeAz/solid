package product

type ProductOffer struct {
	Product Product
	Offer   float64
}

func NewProductOffer(
	product Product,
	offer float64) ProductOffer {
	return ProductOffer{
		Product: product,
		Offer:   offer,
	}
}

func (p ProductOffer) getPriceWithOffer() float64 {
	return p.Product.Price * (1 - p.Offer)
}

func (p ProductOffer) GetPriceWithTax(tax float64) float64 {
	return (p.getPriceWithOffer() * (tax + p.Product.Category.GetTax())) * float64(p.Product.Amount)
}
