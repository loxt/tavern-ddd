package services

import (
	"github.com/google/uuid"
	"github.com/loxt/tavern-ddd/aggregate"
	"testing"
)

func init_products(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "A nice cold beer", 2.99)

	if err != nil {
		t.Fatal(err)
	}

	peanuts, err := aggregate.NewProduct("Peanuts", "A bag of peanuts", 1.99)

	if err != nil {
		t.Fatal(err)
	}

	wine, err := aggregate.NewProduct("Wine", "A nice bottle of wine", 9.99)

	if err != nil {
		t.Fatal(err)
	}

	return []aggregate.Product{beer, peanuts, wine}
}

func TestOrder_NewOrderService(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(WithMemoryCustomerRepository(), WithMemoryProductRepository(products))

	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("Emannuel")

	if err != nil {
		t.Error(err)
	}

	err = os.customers.Create(cust)

	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = os.CreateOrder(cust.GetID(), order)

	if err != nil {
		t.Error(err)
	}
}
