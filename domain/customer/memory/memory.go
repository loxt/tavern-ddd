package memory

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/loxt/tavern-ddd/domain/customer"
	"sync"
)

type MemoryRepository struct {
	customers map[uuid.UUID]customer.Customer
	sync.Mutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]customer.Customer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (customer.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}
	return customer.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryRepository) Create(c customer.Customer) error {
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]customer.Customer)
		mr.Unlock()
	}

	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToCreateCustomer)
	}

	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()

	return nil
}

func (mr *MemoryRepository) Update(c customer.Customer) error {
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrFailedToUpdateCustomer)
	}

	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()

	return nil
}
