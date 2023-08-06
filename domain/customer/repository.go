package customer

import (
	"errors"
	"github.com/google/uuid"
	"github.com/loxt/tavern-ddd/aggregate"
)

var (
	ErrCustomerNotFound       = errors.New("the customer was not found")
	ErrFailedToCreateCustomer = errors.New("failed to create customer")
	ErrFailedToUpdateCustomer = errors.New("failed to update customer")
)

type CustomerRepository interface {
	Get(uuid uuid.UUID) (aggregate.Customer, error)
	Create(aggregate.Customer) error
	Update(aggregate.Customer) error
}
