package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/loxt/tavern-ddd/aggregate"
	"testing"
)

func Test_Tavern(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(WithMongoCustomerRepository(context.Background(), "mongodb://root:docker@localhost:27017"),
		WithMemoryProductRepository(products))

	if err != nil {
		t.Fatal(err)
	}

	tavern, err := NewTavern(WithOrderService(os))

	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("Emannuel")

	if err != nil {
		t.Fatal(err)
	}

	if err = os.customers.Create(cust); err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(cust.GetID(), order)

	if err != nil {
		t.Fatal(err)
	}

}
