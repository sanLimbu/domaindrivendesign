package services

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/sanLimbu/ddd-go/aggregate"
	"github.com/sanLimbu/ddd-go/domain/customer"
	"github.com/sanLimbu/ddd-go/domain/customer/memory"
	"github.com/sanLimbu/ddd-go/domain/customer/mongo"
	"github.com/sanLimbu/ddd-go/domain/product"
	productMemory "github.com/sanLimbu/ddd-go/domain/product/memory"
)

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}
type OrderConfiguration func(os *OrderService) error

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {

	os := &OrderService{}

	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil

}

func WithCustomerRepository(c customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = c
		return nil
	}
}

func WithMongoCustomerRepository(connectionString string) OrderConfiguration {
	return func(os *OrderService) error {

		cr, err := mongo.New(context.Background(), connectionString)
		if err != nil {
			return err
		}

		os.customers = cr
		return nil

	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
		pr := productMemory.New()

		// Add Items to repo
		for _, p := range products {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}
		//os.products = pr
		return nil
	}
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {

	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	var products []aggregate.Product
	var price float64

	for _, id := range productIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}

		products = append(products, p)
		price += p.GetPrice()
	}
	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))

	return price, nil
}
