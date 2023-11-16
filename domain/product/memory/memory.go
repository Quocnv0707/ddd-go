package memory

import (
	"ddd-go/aggregate"
	"ddd-go/domain/product"
	"sync"

	"github.com/google/uuid"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (repo *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product
	for _, product := range repo.products {
		products = append(products, product)
	}
	return products, nil
}

func (repo *MemoryProductRepository) GetById(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := repo.products[id]; ok {
		return product, nil
	}
	return aggregate.Product{}, product.ErrProductNotFound
}

func (repo *MemoryProductRepository) Add(p aggregate.Product) error {
	repo.Lock()
	defer repo.Unlock()

	if _, ok := repo.products[p.GetID()]; ok {
		return product.ErrProductAlreadyExitst
	}
	repo.products[p.GetID()] = p
	return nil
}

func (repo *MemoryProductRepository) Update(p aggregate.Product) error {
	repo.Lock()
	defer repo.Unlock()

	if _, ok := repo.products[p.GetID()]; !ok {
		return product.ErrProductNotFound
	}

	repo.products[p.GetID()] = p
	return nil
}

func (repo *MemoryProductRepository) Delete(id uuid.UUID) error {
	repo.Lock()
	defer repo.Unlock()

	if _, ok := repo.products[id]; !ok {
		return product.ErrProductNotFound
	}
	delete(repo.products, id)
	return nil
}
