package product

import (
	"ddd-go/aggregate"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound      = errors.New("the product was not found")
	ErrProductAlreadyExitst = errors.New("the product already exitst")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetById(id uuid.UUID) (aggregate.Product, error)
	Add(products aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(uuid.UUID) error
}
