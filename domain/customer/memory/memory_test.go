package memory

import (
	"tavern/domain/customer"
	"testing"

	"github.com/google/uuid"
)

func TestMemory_GetMemory(t *testing.T) {
	type testcase struct {
		name      string
		expectErr error
		id        uuid.UUID
	}

	cus, err := customer.NewCustomer("QuocNV")
	if err != nil {
		t.Fatal(err)
	}

	id := cus.GetID()
	//tạo một repository
	repo := MemoryCustomerRepository{
		customers: map[uuid.UUID]customer.Customer{
			id: cus,
		},
	}

	testcases := []testcase{
		{
			name:      "No customer by ID",
			expectErr: customer.ErrCustomerNotFound,
			id:        uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d423"),
		},
		{
			name:      "Get Customer By ID",
			expectErr: nil,
			id:        id,
		},
	}

	for _, i := range testcases {
		t.Run(i.name, func(t *testing.T) {
			_, ok := repo.Get(i.id)
			if ok != i.expectErr {
				t.Errorf("Expect error %v get %v", i.expectErr, err)
			}
		})
	}
}

func TestMemory_AddCustomer(t *testing.T) {
	type testcase struct {
		name      string
		cus       string
		expectErr error
	}

	testcases := []testcase{
		{
			name:      "add customer",
			cus:       "Quoc",
			expectErr: nil,
		},
	}

	for _, i := range testcases {
		t.Run(i.name, func(t *testing.T) {
			repo := MemoryCustomerRepository{
				customers: make(map[uuid.UUID]customer.Customer),
			}
			cust, err := customer.NewCustomer(i.cus)
			if err != nil {
				t.Fatal(err)
			}
			err = repo.Add(cust)
			if err != i.expectErr {
				t.Errorf("Expect error %v, got %v", i.expectErr, err)
			}
			found, ok := repo.Get(cust.GetID())
			if ok != nil {
				t.Fatal(err)
			}
			if found.GetID() != cust.GetID() {
				t.Errorf("Expect error %v, got %v", i.expectErr, err)
			}
		})

	}
}

func TestMemory_UpdateCustomer(t *testing.T) {
	type testcase struct {
		name      string
		cus       string
		id        uuid.UUID
		expectErr error
	}
	cus, err := customer.NewCustomer("QuocNV")
	if err != nil {
		t.Fatal(err)
	}
	id := cus.GetID()
	repo := MemoryCustomerRepository{
		customers: map[uuid.UUID]customer.Customer{
			id: cus,
		},
	}
	testcases := []testcase{
		{
			name:      "Update Customer",
			cus:       "NguyenVanQuoc",
			expectErr: nil,
			id:        id,
		},
	}

	for _, i := range testcases {
		t.Run(i.name, func(t *testing.T) {
			cust, err := customer.NewCustomer(i.cus)
			if err != nil {
				t.Fatal(err)
			}
			err = repo.Update(cust)
			if err != i.expectErr {
				t.Errorf("Expect Error %v, got %v", i.expectErr, err)
			}

		})
	}

}
