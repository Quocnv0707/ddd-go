package customer_test

import (
	"tavern/domain/customer"
	"testing"
)

type testCase struct {
	name      string
	test      string
	expectErr error
}

func TestCutomer_NewCustomer(t *testing.T) {
	testCases := []testCase{
		{
			name:      "",
			test:      "Invalid Person Error",
			expectErr: customer.ErrInvalidPerson,
		}, {
			name:      "Nguyen Van Quoc",
			test:      "valid Person Name",
			expectErr: nil,
		},
	}
	for _, i := range testCases {
		t.Run(i.test, func(t *testing.T) {
			_, err := customer.NewCustomer(i.name)

			if err != i.expectErr {
				t.Errorf("Expect Error %v, got %v", i.expectErr, err)
			}
		})
	}
}
