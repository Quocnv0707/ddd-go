package aggregate

import (
	"ddd-go/entity"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrMissingValue = errors.New("missing value")
)

type Product struct {
	item     *entity.Item
	quantity int
	price    float64
}

func NewProduct(name, descripttion string, price float64) (Product, error) {
	if name == "" || descripttion == "" {
		return Product{}, ErrMissingValue
	}
	return Product{
		item: &entity.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: descripttion,
		},
		price:    price,
		quantity: 0,
	}, nil
}

func (p *Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p *Product) GetItem() *entity.Item {
	return p.item
}

func (p *Product) GetPrice() float64 {
	return p.price
}
