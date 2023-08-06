package services

import (
	"github.com/google/uuid"
	"github.com/loxt/tavern-ddd/domain/customer"
	"github.com/loxt/tavern-ddd/domain/customer/memory"
	"log"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
}

func NewOrderService(configs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	for _, config := range configs {
		if err := config(os); err != nil {
			return nil, err
		}
	}

	return os, nil
}

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.NewMemoryRepository()
	return WithCustomerRepository(cr)
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) error {
	c, err := o.customers.Get(customerID)

	if err != nil {
		return err
	}

	log.Println(c)
	return nil
}
