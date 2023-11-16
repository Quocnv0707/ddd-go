package service

import (
	"context"
	"ddd-go/aggregate"
	"ddd-go/domain/customer"
	"ddd-go/domain/customer/memory"
	"ddd-go/domain/product"
	proMemory "ddd-go/domain/product/memory"
	"log"

	"ddd-go/domain/customer/mongo"

	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	prpducts  product.ProductRepository
}

func NewOrderService(configs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}
	for _, cfg := range configs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
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

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		proRepo := proMemory.New()
		for _, item := range products {
			err := proRepo.Add(item)
			if err != nil {
				return err
			}
		}
		os.prpducts = proRepo
		return nil
	}
}

func (os *OrderService) CreateOrder(customerID uuid.UUID, productsID []uuid.UUID) (float64, error) {
	cus, err := os.customers.Get(customerID)
	if err != nil {
		return 0, err
	}
	var products []aggregate.Product
	var price float64
	for _, id := range productsID {
		p, err := os.prpducts.GetById(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		price += p.GetPrice()
	}
	log.Printf("Customer: has ID: %s has %d products[%v]", cus.GetID(), len(products), products)
	return price, nil
}
