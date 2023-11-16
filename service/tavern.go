package service

import (
	"log"

	"github.com/google/uuid"
)

type TavernConfiguration func(ts *Tavern) error

type Tavern struct {
	OrderService   *OrderService
	BillingService interface{}
}

func NewTavern(cgfs ...TavernConfiguration) (*Tavern, error) {
	t := &Tavern{}
	for _, cfg := range cgfs {
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func WithOrderService(os *OrderService) TavernConfiguration {
	return func(ts *Tavern) error {
		ts.OrderService = os
		return nil
	}
}

func (t *Tavern) Order(customerID uuid.UUID, productsID []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customerID, productsID)
	if err != nil {
		return err
	}
	log.Printf("Bill the customer: %0.0f", price)
	// Bill the customer
	//err = t.BillingService.Bill(customer, price)
	return nil
}
