package customer

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound      = errors.New("the cusomer was not found in this repository")
	ErrFailedToAddCustomer   = errors.New("failed to add customer to the repository")
	ErrFaildToUpdateCustomer = errors.New("failed to update customer in the repository")
)

type CustomerRepository interface {
	Get(uuid.UUID) (Customer, error)
	Add(Customer) error
	Update(Customer) error
}
