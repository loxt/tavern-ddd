package product

import (
	"errors"
	"github.com/google/uuid"
	"github.com/loxt/tavern-ddd/aggregate"
)

var (
	ErrProductNotFound = errors.New("product not found")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (*aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}
