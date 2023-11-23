package mongo

import (
	"context"
	"tavern/domain/customer"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCustomerRepository struct {
	db       *mongo.Database
	customer *mongo.Collection
}
type mongoCustomer struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

func NewFromCustomer(cus customer.Customer) mongoCustomer {
	return mongoCustomer{
		ID:   cus.GetID(),
		Name: cus.GetName(),
	}
}

func (m *mongoCustomer) ToAggregate() customer.Customer {
	cus := customer.Customer{}
	cus.SetID(m.ID)
	cus.SetName(m.Name)
	return cus
}

// Create a new mongodb repository
func NewMongoCustomerRepository(ctx context.Context, connectionString string) (*MongoCustomerRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}
	db := client.Database("ddd")
	customers := db.Collection("customers")

	return &MongoCustomerRepository{
		db:       db,
		customer: customers,
	}, nil
}

func (mongoRepo *MongoCustomerRepository) Get(id uuid.UUID) (customer.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := mongoRepo.customer.FindOne(ctx, bson.M{"id": id})
	var c mongoCustomer
	err := result.Decode(&c)
	if err != nil {
		return customer.Customer{}, err
	}
	return c.ToAggregate(), nil

}

func (mongoRepo *MongoCustomerRepository) Add(cus customer.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()


	internal := NewFromCustomer(cus)
	_, err := mongoRepo.customer.InsertOne(ctx, internal)
	if err != nil {
		return err
	}
	return nil
}

func (mongoRepo *MongoCustomerRepository) Update(cus customer.Customer) error {
	panic("to implement")
}
