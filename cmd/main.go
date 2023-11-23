package main

import (
	"fmt"
	"tavern/config"
	"tavern/domain/product"
	"tavern/service/order"
	servicetavern "tavern/service/tavern"

	"github.com/google/uuid"
)

func main() {
	var connectionString = fmt.Sprintf("mongodb+srv://quocnv07:%s@cluster0.tgzadpv.mongodb.net/?retryWrites=true&w=majority", config.DB_PASS)

	products := productInventory()
	// Create Order Service to use in tavern
	os, err := order.NewOrderService(
		// order.WithMemoryCustomerRepository(),
		order.WithMongoCustomerRepository(connectionString),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		panic(err)
	}
	// Create tavern service
	tavern, err := servicetavern.NewTavern(
		servicetavern.WithOrderService(os))
	if err != nil {
		panic(err)
	}

	uid, err := os.AddCustomer("Percy")
	if err != nil {
		panic(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	// Execute Order
	err = tavern.Order(uid, order)
	if err != nil {
		panic(err)
	}
}

func productInventory() []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		panic(err)
	}
	peenuts, err := product.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		panic(err)
	}
	wine, err := product.NewProduct("Wine", "Healthy Snacks", 0.99)
	if err != nil {
		panic(err)
	}
	products := []product.Product{
		beer, peenuts, wine,
	}
	return products
}
