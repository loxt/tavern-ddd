package main

import (
	"context"
	"github.com/google/uuid"
	"github.com/loxt/tavern-ddd/domain/product"
	"github.com/loxt/tavern-ddd/services/order"
	"github.com/loxt/tavern-ddd/services/tavern"
)

func main() {
	products := productInventory()

	os, err := order.NewOrderService(
		order.WithMongoCustomerRepository(context.Background(), "mongodb://root:docker@localhost:27017"),
		order.WithMemoryProductRepository(products))

	if err != nil {
		panic(err)
	}

	tavern, err := tavern.NewTavern(
		tavern.WithOrderService(os))

	if err != nil {
		panic(err)
	}

	uid, err := os.CreateCustomer("Emannuel")

	if err != nil {
		panic(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(uid, order)

	if err != nil {
		panic(err)
	}
}

func productInventory() []product.Product {
	beer, err := product.NewProduct("Beer", "A nice cold beer", 2.99)

	if err != nil {
		panic(err)
	}

	peanuts, err := product.NewProduct("Peanuts", "A bag of peanuts", 1.99)

	if err != nil {
		panic(err)
	}

	wine, err := product.NewProduct("Wine", "A nice bottle of wine", 9.99)

	if err != nil {
		panic(err)
	}

	return []product.Product{beer, peanuts, wine}
}
