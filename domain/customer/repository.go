package customer

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound       = errors.New("the customer was not found")
	ErrFailedToCreateCustomer = errors.New("failed to create customer")
	ErrFailedToUpdateCustomer = errors.New("failed to update customer")
)

type Repository interface {
	Get(uuid uuid.UUID) (Customer, error)
	Create(Customer) error
	Update(Customer) error
}
