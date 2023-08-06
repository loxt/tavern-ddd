package services

import (
	"github.com/google/uuid"
	"github.com/loxt/tavern-ddd/aggregate"
	"github.com/loxt/tavern-ddd/domain/customer"
	"github.com/loxt/tavern-ddd/domain/customer/memory"
	"github.com/loxt/tavern-ddd/domain/product"
	prodMem "github.com/loxt/tavern-ddd/domain/product/memory"
	"log"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
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

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodMem.NewMemoryProductRepository()

		for _, p := range products {
			if err := pr.Create(p); err != nil {
				return err
			}
		}

		os.products = pr
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

	var products []aggregate.Product
	var total float64

	for _, id := range productsIDs {
		p, err := o.products.GetByID(id)

		if err != nil {
			return err
		}

		products = append(products, p)
		total += p.GetPrice()
	}

	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))
	return nil
}
