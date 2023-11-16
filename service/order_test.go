package service

import (
	"ddd-go/aggregate"
	"testing"

	"github.com/google/uuid"
)

func initProducts(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "4.5%", 16.5)
	if err != nil {
		t.Error(err)
	}
	peanut, err := aggregate.NewProduct("Peanut", "4%", 17.5)
	if err != nil {
		t.Error(err)
	}
	tea, err := aggregate.NewProduct("Tea", "5%", 18.5)
	if err != nil {
		t.Error(err)
	}
	products := []aggregate.Product{
		beer, peanut, tea,
	}
	return products
}
func TestOrder_CreaterOrder(t *testing.T) {
	products := initProducts(t)
	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}
	//add customererr
	cus, err := aggregate.NewCustomer("Hary")
	if err != nil {
		t.Error(err)
	}
	err = os.customers.Add(cus)
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	_, err = os.CreateOrder(cus.GetID(), order)
	if err != nil {
		t.Error(err)
	}

}
