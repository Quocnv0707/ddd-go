package memory

import (
	"ddd-go/aggregate"
	"ddd-go/domain/customer"
	"sync"

	"github.com/google/uuid"
)

type MemoryCustomerRepository struct {
	cusomers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryCustomerRepository {
	return &MemoryCustomerRepository{
		cusomers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (repo *MemoryCustomerRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := repo.cusomers[id]; ok {
		return customer, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (repo *MemoryCustomerRepository) Add(cus aggregate.Customer) error {
	//kiểm tra repo được tạo hay chưa
	if repo.cusomers == nil {
		repo.Lock()
		repo.cusomers = make(map[uuid.UUID]aggregate.Customer)
		repo.Unlock()
	}
	if _, ok := repo.cusomers[cus.GetID()]; ok {
		return customer.ErrFailedToAddCustomer
	}
	repo.Lock()
	repo.cusomers[cus.GetID()] = cus
	repo.Unlock()
	return nil
}

func (repo *MemoryCustomerRepository) Update(cus aggregate.Customer) error {
	if _, ok := repo.cusomers[cus.GetID()]; !ok {
		return customer.ErrFaildToUpdateCustomer
	}
	repo.Lock()
	repo.cusomers[cus.GetID()] = cus
	repo.Unlock()
	return nil
}
