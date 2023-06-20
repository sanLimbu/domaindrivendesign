package aggregate_test

import (
	"testing"

	"github.com/sanLimbu/ddd-go/aggregate"
)

func TestCustomer_NewCustomer(t *testing.T) {

	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		// {
		// 	test:        "Empty name validation",
		// 	name:        "",
		// 	expectedErr: aggregate.ErrInvalidPerson,
		// },
		{
			test:        "Valid Name",
			name:        "santosh limbu",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
