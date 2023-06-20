package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/sanLimbu/ddd-go/aggregate"
)

var (
	// ErrCustomerNotFound is returned when a customer is not found.
	ErrCustomerNotFound = errors.New("the customer was not found in the repository")
	// ErrFailedToAddCustomer is returned when the customer could not be added to the repository.
	ErrFailedToAddCustomer = errors.New("failed to add the customer to the repository")
	// ErrUpdateCustomer is returned when the customer could not be updated in the repository.
	ErrUpdateCustomer = errors.New("failed to update the customer in the repository")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
