package memory

import (
	"tavern/domain/product"
	"testing"

	"github.com/google/uuid"
)

func TestMemory_AddProduct(t *testing.T) {
	repo := New()
	p, err := product.NewProduct("Hala", "HoaKy", 2.5)
	if err != nil {
		t.Error(err)
	}
	repo.Add(p)
	if len(repo.products) != 1 {
		t.Errorf("Expect 1 Product , got %d", len(repo.products))
	}

}

func TestMemory_GetAll(t *testing.T) {
	repo := New()
	for i := 0; i < 3; i++ {
		// name, desc := (string) i
		p, err := product.NewProduct("Pencil", "Long: 6cm", 6.3)
		if err != nil {
			t.Error(err)
		}
		repo.Add(p)
	}
	list, _ := repo.GetAll()
	if len(list) != 3 {
		t.Errorf("Expect 3 product, got %d", len(list))
	}
}

func TestMemory_GetProductByID(t *testing.T) {
	repo := New()
	p, err := product.NewProduct("Nho", "Hoa ky", 9.9)
	if err != nil {
		t.Error(err)
	}
	repo.Add(p)
	if len(repo.products) <= 0 {
		t.Errorf("Expect 1 product, go %v", len(repo.products))
	}

	type testCase struct {
		name      string
		id        uuid.UUID
		expectErr error
	}

	testCases := []testCase{
		{
			name:      "Get product by ID",
			id:        p.GetID(),
			expectErr: nil,
		},
		{
			name:      "Get non-existing product by id",
			id:        uuid.New(),
			expectErr: product.ErrProductNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.GetById(tc.id)
			if err != tc.expectErr {
				t.Errorf("experct error %v, get %v", tc.expectErr, err)
			}
		})
	}
}

func TestMemory_DeleteProduct(t *testing.T) {
	repo := New()
	p, err := product.NewProduct("bread", "made from wheat flour", 4.0)
	if err != nil {
		t.Error(err)
	}
	err = repo.Add(p)
	if err != nil {
		t.Error(err)
	}
	err = repo.Delete(p.GetID())
	if err != nil {
		t.Error(err)
	}
	if len(repo.products) != 0 {
		t.Errorf("Expect 0 products, got %v", len(repo.products))
	}
}
