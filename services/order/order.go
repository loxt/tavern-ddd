package order

import (
	"context"
	"github.com/google/uuid"
	"github.com/loxt/tavern-ddd/domain/customer"
	"github.com/loxt/tavern-ddd/domain/customer/memory"
	"github.com/loxt/tavern-ddd/domain/customer/mongo"
	"github.com/loxt/tavern-ddd/domain/product"
	prodMem "github.com/loxt/tavern-ddd/domain/product/memory"
	"log"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.Repository
	products  product.Repository
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

func WithCustomerRepository(cr customer.Repository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryProductRepository(products []product.Product) OrderConfiguration {
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

func WithMongoCustomerRepository(ctx context.Context, connStr string) OrderConfiguration {
	return func(os *OrderService) error {
		cr, err := mongo.NewMongoRepository(ctx, connStr)

		if err != nil {
			return err
		}

		os.customers = cr
		return nil
	}
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) (float64, error) {
	c, err := o.customers.Get(customerID)

	if err != nil {
		return 0, err
	}

	var products []product.Product
	var total float64

	for _, id := range productsIDs {
		p, err := o.products.GetByID(id)

		if err != nil {
			return 0, err
		}

		products = append(products, p)
		total += p.GetPrice()
	}

	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))
	return total, nil
}

func (o *OrderService) CreateCustomer(name string) (uuid.UUID, error) {
	c, err := customer.NewCustomer(name)

	if err != nil {
		return uuid.Nil, err
	}

	err = o.customers.Create(c)

	if err != nil {
		return uuid.Nil, err
	}

	return c.GetID(), nil
}
