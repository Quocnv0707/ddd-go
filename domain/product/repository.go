package product

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound      = errors.New("the product was not found")
	ErrProductAlreadyExitst = errors.New("the product already exitst")
)

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetById(id uuid.UUID) (Product, error)
	Add(products Product) error
	Update(product Product) error
	Delete(uuid.UUID) error
}
