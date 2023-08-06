package product

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrProductNotFound      = errors.New("product not found")
	ErrProductAlreadyExists = errors.New("product already exists")
)

type Repository interface {
	GetAll() ([]Product, error)
	GetByID(id uuid.UUID) (Product, error)
	Create(product Product) error
	Update(product Product) error
	Delete(id uuid.UUID) error
}
