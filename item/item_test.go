package item

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/sahaj279/go-assignment/item/enum"
)

type InputItem struct {
	name     string
	itemType string
	price    float64
	quantity int
}

func TestCreateItem(t *testing.T) {
	var itemSvc ItemSvc
	tests := []struct {
		scenario     string
		expectedItem Item
		input        InputItem
		err          error
	}{
		{
			scenario:     "Valid test case",
			expectedItem: Item{Name: "bread", Type: enum.Raw, Price: 30, Quantity: 3},
			input:        InputItem{"bread", "rAw", 30, 3},
			err:          nil,
		},
		{
			scenario:     "InValid test case invalid item type",
			expectedItem: Item{Name: "bread", Type: enum.Raw, Price: 30, Quantity: 3},
			input:        InputItem{"bread", "ram", 30, 3},
			err:          errors.New("CreateItem: ram does not belong to ItemType values"),
		},
		{
			scenario:     "InValid test case invalid price and quantity",
			expectedItem: Item{Name: "bread", Type: enum.Raw, Price: 30, Quantity: 3},
			input:        InputItem{"bread", "raw", -30, -3},
			err:          errors.New(" CreateItem: Price: validate: negative value; Quantity: validate: negative value."),
		},
	}

	for _, tc := range tests {
		result, err := itemSvc.CreateItem(tc.input.name, tc.input.itemType, tc.input.price, tc.input.quantity)
		if err != nil && tc.err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
		if result != tc.expectedItem && tc.err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}

func TestCalculateTex(t *testing.T) {
	var itemSvc ItemSvc
	tests := []struct {
		scenario string
		item     Item
		tax      float64
	}{
		{
			scenario: "tax calculation for raw",
			item: Item{
				Name:  "Bread",
				Type:  enum.Raw,
				Price: 100,
			},
			tax: 12.5,
		},
		{
			scenario: "tax calculation for manufactured",
			item: Item{
				Name:  "Bread",
				Type:  enum.Manufactured,
				Price: 100,
			},
			tax: 14.75,
		},
		{
			scenario: "tax calculation for imported final cost up to 100",
			item: Item{
				Name:  "Bread",
				Type:  enum.Imported,
				Price: 10,
			},
			tax: 6,
		},
		{
			scenario: "final price calculation for imported final cost >100 and <= 200",
			item: Item{
				Name:  "Bread",
				Type:  enum.Imported,
				Price: 100,
			},
			tax: 20,
		},
		{
			scenario: "final price calculation for imported final cost > 200",
			item: Item{
				Name:  "Bread",
				Type:  enum.Imported,
				Price: 200,
			},
			tax: 31,
		},
	}

	for _, tc := range tests {
		tax := itemSvc.CalculateTax(&tc.item)
		if tc.tax != tax {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, tax, tc.tax)
		}
	}
}
