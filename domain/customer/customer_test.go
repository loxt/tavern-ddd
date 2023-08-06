package customer_test

import (
	"errors"
	"github.com/loxt/tavern-ddd/domain/customer"
	"testing"
)

func TestNewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty name validation",
			name:        "",
			expectedErr: customer.ErrInvalidPerson,
		},
		{
			test:        "Valid name",
			name:        "Emannuel Matos",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := customer.NewCustomer(tc.name)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
