package product_test

import (
	"tavern/domain/product"
	"testing"
)

func TestProduct_NewProduct(t *testing.T) {
	type testCase struct {
		test      string
		name      string
		desc      string
		price     float64
		expectErr error
	}

	testCases := []testCase{
		{
			test:      "invalid value: empty Name",
			name:      "",
			expectErr: product.ErrMissingValue,
		},
		{
			test:      "valid value",
			name:      "Golem",
			desc:      "Made in Vietnam",
			price:     650000.5,
			expectErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := product.NewProduct(tc.name, tc.desc, tc.price)
			if err != tc.expectErr {
				t.Errorf("Expect Error %v, got %v", tc.expectErr, err)
			}

		})
	}
}
