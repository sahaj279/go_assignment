package item

import (
	"assignment1/item/enum"
	"testing"
)

func TestCreateItem(t *testing.T) {
	correctItem := Item{Name: "bread", Type: enum.Raw, Price: 40, Quantity: 100}
	result, _ := CreateItem("bread", "raw", 40, 100)
	if result != correctItem {
		t.Errorf("item Creation failed")
	}
}
func TestCalculateSalesTax(t *testing.T) {
	var tests = []struct {
		scenario string
		req      Item
		tax      float64
	}{
		{
			scenario: "tax calculation for raw",
			req: Item{
				Name:  "Bread",
				Type:  enum.Raw,
				Price: 100,
			},
			tax: 12.5,
		},
		{
			scenario: "tax calculation for manufactured",
			req: Item{
				Name:  "Bread",
				Type:  enum.Manufactured,
				Price: 100,
			},
			tax: 14.75,
		},
		{
			scenario: "tax calculation for imported final cost up to 100",
			req: Item{
				Name:  "Bread",
				Type:  enum.Imported,
				Price: 10,
			},
			tax: 6,
		},
		{
			scenario: "final price calculation for imported final cost >100 and <= 200",
			req: Item{
				Name:  "Bread",
				Type:  enum.Imported,
				Price: 100,
			},
			tax: 20,
		},
		{
			scenario: "final price calculation for imported final cost > 200",
			req: Item{
				Name:  "Bread",
				Type:  enum.Imported,
				Price: 200,
			},
			tax: 31,
		},
	}

	for _, tc := range tests {
		tc.req.CalculateSalesTax()
		if tc.tax != tc.req.SalesTax {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, tc.req.SalesTax, tc.tax)
		}
	}
}
