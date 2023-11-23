package service

import (
	"fmt"
	"tavern/config"
	"tavern/domain/product"
	"tavern/service/order"
	"testing"

	"github.com/google/uuid"
)

func initProducts(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "4.5%", 16.5)
	if err != nil {
		t.Error(err)
	}
	peanut, err := product.NewProduct("Peanut", "4%", 17.5)
	if err != nil {
		t.Error(err)
	}
	tea, err := product.NewProduct("Tea", "5%", 18.5)
	if err != nil {
		t.Error(err)
	}
	products := []product.Product{
		beer, peanut, tea,
	}
	return products
}

func TestTavern(t *testing.T) {
	products := initProducts(t)
	os, err := order.NewOrderService(
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
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
	cusID, err := os.AddCustomer("QuocNV")
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	err = tavern.Order(cusID, order)
	if err != nil {
		t.Error(err)
	}

}

func Test_MongoTavern(t *testing.T) {
	products := initProducts(t)
	var connectionString = fmt.Sprintf("mongodb+srv://quocnv07:%s@cluster0.tgzadpv.mongodb.net/?retryWrites=true&w=majority", config.DB_PASS) 
	os, err := order.NewOrderService(
		order.WithMongoCustomerRepository(connectionString),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}
	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}
	cusID, err := os.AddCustomer("QUOC")
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
		products[1].GetID(),
	}
	err = tavern.Order(cusID, order)
	if err != nil {
		t.Error(err)
	}
}
