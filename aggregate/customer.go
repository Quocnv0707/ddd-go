package aggregate

import (
	"ddd-go/entity"
	"ddd-go/valueoject"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidPerson = errors.New("Customer must have a valid person")
)

type Customer struct {
	person   *entity.Person
	products []*entity.Item

	transactions []*valueoject.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}
	person := &entity.Person{
		ID:   uuid.New(),
		Name: name,
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]*valueoject.Transaction, 0),
	}, nil
}

func (cus *Customer) GetID() uuid.UUID {
	return cus.person.ID
}

func (cus *Customer) SetID(id uuid.UUID) {
	if cus.person == nil {
		cus.person = &entity.Person{}
	}
	cus.person.ID = id
}

func (cus *Customer) GetName() string {
	return cus.person.Name
}

func (cus *Customer) SetName(name string) {
	if cus.person == nil {
		cus.person = &entity.Person{}
	}
	cus.person.Name = name
}
