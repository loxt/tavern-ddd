package tavern

import (
	"context"
	"github.com/google/uuid"
	"github.com/loxt/tavern-ddd/domain/product"
	order "github.com/loxt/tavern-ddd/services/order"
	"testing"
)

func Test_Tavern(t *testing.T) {
	products := init_products(t)

	os, err := order.NewOrderService(order.WithMongoCustomerRepository(context.Background(), "mongodb://root:docker@localhost:27017"),
		order.WithMemoryProductRepository(products))

	if err != nil {
		t.Fatal(err)
	}

	tavern, err := NewTavern(WithOrderService(os))

	if err != nil {
		t.Fatal(err)
	}

	uid, err := os.CreateCustomer("Emannuel")

	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(uid, order)

	if err != nil {
		t.Fatal(err)
	}
}

func init_products(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "A nice cold beer", 2.99)

	if err != nil {
		t.Fatal(err)
	}

	peanuts, err := product.NewProduct("Peanuts", "A bag of peanuts", 1.99)

	if err != nil {
		t.Fatal(err)
	}

	wine, err := product.NewProduct("Wine", "A nice bottle of wine", 9.99)

	if err != nil {
		t.Fatal(err)
	}

	return []product.Product{beer, peanuts, wine}
}
