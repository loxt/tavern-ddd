package product

import (
	"errors"
	"github.com/google/uuid"
	"github.com/loxt/tavern-ddd/domain"
)

type Product struct {
	item     *tavern.Item
	price    float64
	quantity int
}

var (
	ErrMissingValues = errors.New("missing important values")
)

func NewProduct(name string, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrMissingValues
	}

	return Product{
		item: &tavern.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 0,
	}, nil
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() *tavern.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}

func (p Product) GetQuantity() int {
	return p.quantity
}
