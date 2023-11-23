package service

import (
	"log"
	"tavern/service/order"

	"github.com/google/uuid"
)

type TavernConfiguration func(ts *Tavern) error

type Tavern struct {
	OrderService   *order.OrderService
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

func WithOrderService(os *order.OrderService) TavernConfiguration {
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
	log.Printf("Bill the customer: %0.0f dolar", price)
	// Bill the customer
	//err = t.BillingService.Bill(customer, price)
	return nil
}
