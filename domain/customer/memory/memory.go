package memory

import (
	"sync"
	"tavern/domain/customer"

	"github.com/google/uuid"
)

type MemoryCustomerRepository struct {
	customers map[uuid.UUID]customer.Customer
	sync.Mutex
}

func New() *MemoryCustomerRepository {
	return &MemoryCustomerRepository{
		customers: make(map[uuid.UUID]customer.Customer),
	}
}

func (repo *MemoryCustomerRepository) Get(id uuid.UUID) (customer.Customer, error) {
	if customer, ok := repo.customers[id]; ok {
		return customer, nil
	}
	return customer.Customer{}, customer.ErrCustomerNotFound
}

func (repo *MemoryCustomerRepository) Add(cus customer.Customer) error {
	//kiểm tra repo được tạo hay chưa
	if repo.customers == nil {
		repo.Lock()
		repo.customers = make(map[uuid.UUID]customer.Customer)
		repo.Unlock()
	}
	if _, ok := repo.customers[cus.GetID()]; ok {
		return customer.ErrFailedToAddCustomer
	}
	repo.Lock()
	repo.customers[cus.GetID()] = cus
	repo.Unlock()
	return nil
}

func (repo *MemoryCustomerRepository) Update(cus customer.Customer) error {
	if _, ok := repo.customers[cus.GetID()]; !ok {
		return customer.ErrFaildToUpdateCustomer
	}
	repo.Lock()
	repo.customers[cus.GetID()] = cus
	repo.Unlock()
	return nil
}
