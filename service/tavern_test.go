package service

import (
	"ddd-go/aggregate"
	"testing"

	"github.com/google/uuid"
)

func TestTavern(t *testing.T) {
	products := initProducts(t)
	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}
	tavern, err := NewTavern(
		WithOrderService(os),
	)
	if err != nil {
		t.Error(err)
	}
	cust, err := aggregate.NewCustomer("QuocNV")
	if err != nil {
		t.Error(err)
	}
	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	err = tavern.Order(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}

}

func Test_MongoTavern(t *testing.T) {
	products := initProducts(t)

	os, err := NewOrderService(
		WithMongoCustomerRepository("mongodb://localhost:27017"),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}
	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	cus, err := aggregate.NewCustomer("Quov")
	if err != nil {
		t.Error(err)
	}
	err = os.customers.Add(cus)
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
		products[1].GetID(),
	}
	err = tavern.Order(cus.GetID(), order)
	if err != nil {
		t.Error(err)
	}
}
