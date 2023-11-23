package memory

import (
	"sync"
	"tavern/domain/product"

	"github.com/google/uuid"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]product.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]product.Product),
	}
}

func (repo *MemoryProductRepository) GetAll() ([]product.Product, error) {
	var products []product.Product
	for _, product := range repo.products {
		products = append(products, product)
	}
	return products, nil
}

func (repo *MemoryProductRepository) GetById(id uuid.UUID) (product.Product, error) {
	if product, ok := repo.products[id]; ok {
		return product, nil
	}
	return product.Product{}, product.ErrProductNotFound
}

func (repo *MemoryProductRepository) Add(p product.Product) error {
	repo.Lock()
	defer repo.Unlock()

	if _, ok := repo.products[p.GetID()]; ok {
		return product.ErrProductAlreadyExitst
	}
	repo.products[p.GetID()] = p
	return nil
}

func (repo *MemoryProductRepository) Update(p product.Product) error {
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
