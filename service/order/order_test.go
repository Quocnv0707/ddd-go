package order

import (
	"fmt"
	"tavern/config"
	"tavern/domain/customer"
	"tavern/domain/product"
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
	cus, err := customer.NewCustomer("Hary")
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
func TestOrder_NewOrderWithMongo(t *testing.T){
	var connectionString = fmt.Sprintf("mongodb+srv://quocnv07:%s@cluster0.tgzadpv.mongodb.net/?retryWrites=true&w=majority", config.DB_PASS) 
	products := initProducts(t)
	mos, err := NewOrderService(
		WithMongoCustomerRepository(connectionString),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}
	cusID, err := mos.AddCustomer("Quoc")
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	_, err = mos.CreateOrder(cusID, order)
	if err != nil {
		t.Error(err)
	}
}
