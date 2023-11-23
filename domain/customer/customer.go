package customer

import (
	"errors"
	"tavern"

	"github.com/google/uuid"
)

var (
	ErrInvalidPerson = errors.New("Customer must have a valid person")
)

type Customer struct {
	person   *tavern.Person
	products []*tavern.Item

	transactions []*tavern.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}
	person := &tavern.Person{
		ID:   uuid.New(),
		Name: name,
	}

	return Customer{
		person:       person,
		products:     make([]*tavern.Item, 0),
		transactions: make([]*tavern.Transaction, 0),
	}, nil
}

func (cus *Customer) GetID() uuid.UUID {
	return cus.person.ID
}

func (cus *Customer) SetID(id uuid.UUID) {
	if cus.person == nil {
		cus.person = &tavern.Person{}
	}
	cus.person.ID = id
}

func (cus *Customer) GetName() string {
	return cus.person.Name
}

func (cus *Customer) SetName(name string) {
	if cus.person == nil {
		cus.person = &tavern.Person{}
	}
	cus.person.Name = name
}
