package service

import (
	"testing"

	"github.com/pkg/errors"
)

func TestNew(t *testing.T) {
	tests := []struct {
		scenario string
		name     string
		price    float64
		quantity int
		typeItem string
		err      error
	}{
		{
			scenario: "all item details provided",
			name:     "Mango",
			price:    100,
			quantity: 2,
			typeItem: "RAw",
			err:      nil,
		},
		{
			scenario: "all item details provided",
			name:     "Orange",
			price:    100,
			quantity: 2,
			typeItem: "imported",
			err:      nil,
		},
		{
			scenario: "quantity less than 0",
			name:     "Orange",
			price:    100,
			quantity: -2,
			typeItem: "imported",
			err:      errors.New("negative value"),
		},
		{
			scenario: "type of item not matches predefined type",
			name:     "Mango",
			price:    100,
			quantity: 2,
			typeItem: "exported",
			err:      errors.New("invalid item type"),
		},
		{
			scenario: "item type not provided",
			name:     "Mango",
			price:    100,
			quantity: 2,
			err:      errors.New("negative value"),
		},
		{
			scenario: "quantity is not provided and mandatory field(item type) is provided",
			name:     "Mango",
			price:    100,
			typeItem: "raw",
			err:      nil,
		},
		{
			scenario: "price is not provided and mandatory field(item type) is provided",
			name:     "Mango",
			quantity: 2,
			typeItem: "raw",
			err:      nil,
		},
		{
			scenario: "name is not provided and mandatory field(item type) is provided",
			price:    100,
			quantity: 2,
			typeItem: "raw",
			err:      nil,
		},
		{
			scenario: "price less than zero",
			price:    -100,
			quantity: 2,
			typeItem: "raw",
			err:      errors.New("negative value"),
		},
	}

	for _, tc := range tests {
		_, err := NewItem(tc.name, tc.typeItem, tc.price, tc.quantity)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}

func TestGetTax(t *testing.T) {
	tests := []struct {
		scenario string
		req      Item
		tax      float64
		err      error
	}{
		{
			scenario: "tax calculation for raw",
			req: Item{
				Name:     "Mango",
				Price:    12,
				Quantity: 1,
				Type:     0,
			},
			tax: 1.5,
			err: nil,
		},
		{
			scenario: "tax calculation for manufactured",
			req: Item{
				Name:     "Mango",
				Price:    12,
				Quantity: 1,
				Type:     1,
			},
			tax: 1.77,
			err: nil,
		},
		{
			scenario: "tax calculation for imported",
			req: Item{
				Name:     "Mango",
				Price:    12.0,
				Quantity: 1,
				Type:     2,
			},
			tax: float64(0.1) * 12,
			err: nil,
		},
		{
			scenario: "tax calculation for imported",
			req: Item{
				Name:     "Orange",
				Price:    100,
				Quantity: 1,
				Type:     2,
			},
			tax: 10,
			err: nil,
		},
		{
			scenario: "tax calculation for imported",
			req: Item{
				Name:     "Tomato",
				Price:    1000,
				Quantity: 1,
				Type:     2,
			},
			tax: 100,
			err: nil,
		},
	}

	for _, tc := range tests {
		tax := tc.req.getTax()
		if tc.tax != tax {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, tax, tc.tax)
		}
	}
}

func TestGetEffectivePrice(t *testing.T) {
	tests := []struct {
		scenario string
		req      Item
		tax      float64
		price    float64
		err      error
	}{
		{
			scenario: "final price calculation for raw",
			req: Item{
				Name:     "Mango",
				Price:    12,
				Quantity: 1,
				Type:     0,
			},
			tax:   1.5,
			price: 13.5,
			err:   nil,
		},
		{
			scenario: "final price calculation for manufactured",
			req: Item{
				Name:     "Mango",
				Price:    12,
				Quantity: 1,
				Type:     1,
			},
			tax:   1.77,
			price: 13.77,
			err:   nil,
		},
		{
			scenario: "final price calculation for imported ( under ImportDutyLimit1)",
			req: Item{
				Name:     "Mango",
				Price:    12.0,
				Quantity: 1,
				Type:     2,
			},
			tax:   float64(0.1) * 12,
			price: 18.2,
			err:   nil,
		},
		{
			scenario: "final price calculation for imported (under ImportDutyLimit2)",
			req: Item{
				Name:     "Orange",
				Price:    100,
				Quantity: 1,
				Type:     2,
			},
			tax:   10,
			price: 120,
			err:   nil,
		},
		{
			scenario: "final price calculation for imported (ExceedeImportDutyLimit2)",
			req: Item{
				Name:     "Tomato",
				Price:    1000,
				Quantity: 1,
				Type:     2,
			},
			tax:   100,
			price: 1155,
			err:   nil,
		},
	}

	for _, tc := range tests {
		price := tc.req.getEffectivePrice()
		if tc.price != price {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, price, tc.price)
		}
	}
}
